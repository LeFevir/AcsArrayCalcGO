// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

func (field *Field) NumericCalcPressureField(procNumber int, array *Array, medium *Medium, numberOfTrisOnLine int, printStat bool) {

	if procNumber <= 0 {
		procNumber = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(procNumber)

	type Empty interface{}
	var empty Empty
	ready := make(chan Empty, field.Nz)

	startTime := time.Now()
	tick := time.Tick(printStatTime * time.Second)

	fmt.Println("Numberic field calculation has started on")
	log.Println("Численный расчет поля начался")

	tris := array.sliceOnTrisProper(numberOfTrisOnLine)

	waveNumber := 2 * math.Pi * array.Frequency / medium.SpeedOfSound
	var p complex128

	for k, zk := range field.Z {
		//Strarting Nz goroutines
		go func(k int, zk float64) {
			for j, yj := range field.Y {
				for i, xi := range field.X {
					p = sumUpRaleigh(tris, NewPoint(xi, yj, zk), waveNumber)
					field.Put(p, i, j, k)
					ready <- empty
				}
			}
		}(k, zk)
	}

	count := field.Nz * field.Ny * field.Nx
	//	Waiting for all goroutines
	for i := 0; i < count; i++ {
		select {
		case <-ready:
		case <-tick:
			if printStat {
				PrintStatus(100.0*i/count, startTime)
			}
		}
	}

	fmt.Printf("Numeric field calculation has finished. It took %s\n", time.Since(startTime))
	log.Printf("Численный расчет поля закончен. Он занял %s\n", time.Since(startTime))
}

func (array *Array) sliceOnTrisProper(numberOfTrisOnLine int) (tris []triangle) {

	numberOfRectsOnLine := numberOfTrisOnLine / 2
	dX := array.Aperture / float64(numberOfRectsOnLine)
	dY := -dX
	a := array.Aperture / 2.0
	for j := 0; j < numberOfRectsOnLine; j++ {
		for i := 0; i < numberOfRectsOnLine; i++ {
			tri1 := NewTriangle(-a+float64(i)*dX, a+float64(j)*dY, -a+float64(i)*dX, a+float64(j+1)*dY, -a+float64(i+1)*dX, a+float64(j)*dY)
			tri2 := NewTriangle(-a+float64(i+1)*dX, a+float64(j)*dY, -a+float64(i)*dX, a+float64(j+1)*dY, -a+float64(i+1)*dX, a+float64(j+1)*dY)

			tri1.computeCentroids()
			if tri1.isNeed(array) {
				tri1.projectOnArray(array.CurvatureRadius)
				tri1.computeCentroids()
				tri1.computeArea()
				tris = append(tris, tri1)
			}

			tri2.computeCentroids()
			if tri2.isNeed(array) {
				tri2.projectOnArray(array.CurvatureRadius)
				tri2.computeCentroids()
				tri2.computeArea()
				tris = append(tris, tri2)
			}
		}
	}
	fmt.Println("Slicing array on triangles has finished. Number of active triangles", len(tris))
	log.Println("Разбиение решетки на треугольники закончено. Активных треугольников", len(tris))
	return
}

func (tri *triangle) projectOnArray(arrayCurvatureRadius float64) {
	F := arrayCurvatureRadius

	x0 := tri.V1.X
	y0 := tri.V1.Y
	tri.V1.Z = F - math.Sqrt(F*F-x0*x0-y0*y0)

	x0 = tri.V2.X
	y0 = tri.V2.Y
	tri.V2.Z = F - math.Sqrt(F*F-x0*x0-y0*y0)

	x0 = tri.V3.X
	y0 = tri.V3.Y
	tri.V3.Z = F - math.Sqrt(F*F-x0*x0-y0*y0)
}

func (tri *triangle) computeCentroids() {
	xc := (tri.V1.X + tri.V2.X + tri.V3.X) / 3.0
	yc := (tri.V1.Y + tri.V2.Y + tri.V3.Y) / 3.0
	zc := (tri.V1.Z + tri.V2.Z + tri.V3.Z) / 3.0
	tri.Center = NewPoint(xc, yc, zc)
}

func (tri *triangle) computeArea() {
	//По теореме Герона
	//TODO Может можно быстрее?
	/*
		Косое произведение векторов
		http://habrahabr.ru/post/147691/
		или вики

		S = abs( (x2-x1)(y3-y1) — (x3-x1)(y2-y1) ) / 2

	*/

	a := math.Sqrt(math.Pow(tri.V1.X-tri.V2.X, 2) + math.Pow(tri.V1.Y-tri.V2.Y, 2) + math.Pow(tri.V1.Z-tri.V2.Z, 2))
	b := math.Sqrt(math.Pow(tri.V1.X-tri.V3.X, 2) + math.Pow(tri.V1.Y-tri.V3.Y, 2) + math.Pow(tri.V1.Z-tri.V3.Z, 2))
	c := math.Sqrt(math.Pow(tri.V2.X-tri.V3.X, 2) + math.Pow(tri.V2.Y-tri.V3.Y, 2) + math.Pow(tri.V2.Z-tri.V3.Z, 2))
	p := 0.5 * (a + b + c)
	tri.Area = math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func (tri *triangle) isNeed(array *Array) bool {
	R2 := array.ElementRadius * array.ElementRadius

	for _, element := range array.Elements {
		// Переходим в лок систему координат каждого элемента
		x0 := tri.Center.X - element.Center.X
		y0 := tri.Center.Y - element.Center.Y

		// Расстояние от центра треугольника до центра элемента
		hypot2 := x0*x0 + y0*y0

		//Сравниваем Расстояние от центра треугольника до центра элемента просто с радиусом элемента
		if hypot2 < R2 {
			cosGamma := math.Cos(math.Asin(math.Hypot(element.Center.X, element.Center.Y) / array.CurvatureRadius))
			a2 := R2
			b2 := cosGamma * cosGamma * R2

			//Сравниваем Расстояние от центра треугольника до центра элемента с наименьшим радиусом эллипса (проекция повернутого элемента на плоскость)
			if hypot2 < b2 {
				return true
			} else {
				phi := math.Atan2(element.Center.Y, element.Center.X)
				beta := math.Atan2(y0, x0)
				rad2 := a2*math.Sin(phi-beta)*math.Sin(phi-beta) + b2*math.Cos(phi-beta)*math.Cos(phi-beta)

				//Сравниваем Расстояние от центра треугольника до центра элемента с остальной частью эллипса (проекция повернутого элемента на плоскость)
				if hypot2 < rad2 {
					return true
				}
			}
		}
	}
	return false
}

func sumUpRaleigh(tris []triangle, point Point, waveNumber float64) (p complex128) {
	for _, tri := range tris {
		r := calculateDistance(tri.Center, point)
		sum := cmplx.Exp(complex(0.0, waveNumber*r))
		sum = sum * complex(tri.Area/r, 0.0)
		p += sum
	}
	p = p * complex(0.0, -waveNumber/(2*math.Pi))
	return
}
