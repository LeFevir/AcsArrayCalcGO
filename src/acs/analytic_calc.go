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

//TODO Избавиться от этой константы
const (
	printStatTime = 10 //seconds for 1 out
)

func (field *Field) AnalytCalcPressureField(procNumber int, array *Array, medium *Medium, printStat bool) {

	if procNumber <= 0 {
		procNumber = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(procNumber)
	fmt.Println("Field calculation has started on", procNumber, "threads")
	log.Println("Начался расчет поля на", procNumber, "потоках")

	type Empty interface{}
	var empty Empty
	ready := make(chan Empty, field.Nz)

	startTime := time.Now()
	tick := time.Tick(printStatTime * time.Second)

	for k, zk := range field.Z {
		//Strarting Nz goroutines
		go func(k int, zk float64) {

			for j, yj := range field.Y {
				for i, xi := range field.X {
					p := AnalytCalcElementsField(array, medium.SpeedOfSound, NewPoint(xi, yj, zk))
					field.Put(p, i, j, k)
				}
				ready <- empty
			}
		}(k, zk)
	}

	count := field.Nz * field.Ny
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

	fmt.Printf("Field calculation has finished. It took %s\n", time.Since(startTime))
	log.Printf("Расчет поля закончен. Он занял %s\n", time.Since(startTime))
}

//Calculates field in point from all active Elements of array
func AnalytCalcElementsField(array *Array, speedOfSound float64, fieldPosition Point) (p complex128) {
	for _, element := range array.Elements {
		waveNumber := 2 * math.Pi * array.Frequency / speedOfSound
		r := calculateDistance(element.Center, fieldPosition)
		theta := calculateDegree(&fieldPosition, &element)
		out := circTransducerFarField(array.ElementRadius, waveNumber, theta, r)

		//Adding phase shift exp(-ikz)
		out = out * cmplx.Exp(complex(0.0, (-waveNumber*element.PhaseShift)))

		p += out
	}
	return
}

//This function calculates far field of plane circular transducer.
//
//           -i*Zr*exp(i*k*r)  2*J1(k*a*sin(theta))
//    p/p0 = ---------------- ---------------------
//                 r             k*a*sin(theta)
//
//Here Zr - rayleigh lenght, k - wavenumber, a - transducer's radius,
//r - distance from transducer's center to the field point,
//theta - angle between transducer's axis and field point
func circTransducerFarField(transRadius, waveNumber, theta, r float64) complex128 {
	// compl = -i*exp(ikr) = -i * (cos(kr) +i *sin(kr)) = -i*cos(kr) + sin(kr)
	//compl := complex(0, -1) * cmplx.Exp(complex(0.0, waveNumber*r))
	compl := complex(math.Sin(waveNumber*r), -math.Cos(waveNumber*r))
	rayleighLenght := waveNumber * transRadius * transRadius / 2
	out := rayleighLenght / r
	// Обрабатываем асимптоту: 2*J1(0)/0 -> 1
	val := math.Sin(theta) * waveNumber * transRadius
	if val != 0.0 {
		out = out * (2 * math.J1(val)) / (val)
	}
	return compl * complex(out, 0.0)
}

//Calculates distance between two points
func calculateDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2) + math.Pow((p1.Z-p2.Z), 2))
}

func calculateDegree(fieldPosition *Point, element *Element) float64 {

	localFieldPosition := calculateLocalFieldPosition(fieldPosition, element)

	return math.Atan(math.Hypot(localFieldPosition.X, localFieldPosition.Y) / localFieldPosition.Z)
}

//Transpose global coordinates of element and field point into local coordinates related to element
func calculateLocalFieldPosition(fieldPosition *Point, element *Element) Point {
	// Смещение осей
	x1 := fieldPosition.X - element.Center.X
	y1 := fieldPosition.Y - element.Center.Y
	z1 := fieldPosition.Z - element.Center.Z

	// Применение матрицы поворота
	x2 := x1*element.Cos_Phi - y1*element.Sin_Phi
	y2 := x1*element.Cos_Gamma*element.Sin_Phi + y1*element.Cos_Gamma*element.Cos_Phi + z1*element.Sin_Gamma
	z2 := -x1*element.Sin_Gamma*element.Sin_Phi - y1*element.Sin_Gamma*element.Cos_Phi + z1*element.Cos_Gamma

	return NewPoint(x2, y2, z2)
}
