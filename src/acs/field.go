// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"log"
	"math/cmplx"
)

const (
	// Множитель для сохранения точности вычисления вещественных чисел в определённых пределах.
	// При выборе множителя = 1000000 можно работать с точностью координат сетки до 1/1000000 метра, то есть до микрометра.
	FLOAT_ACCURACY_MLTPL = 1000000
)

type Field struct {
	Nx int
	Ny int
	Nz int

	MinX float64
	MinY float64
	MinZ float64

	Dx float64
	Dy float64
	Dz float64

	X []float64
	Y []float64
	Z []float64

	Value []complex128

	// Местоположение электронного фокуса
	FocusI int
	FocusJ int
	FocusK int

	// Местоположение максимума (не всегда совпадает с электронным фокусом)
	MainMaxValue float64
	MainMaxI     int
	MainMaxJ     int
	MainMaxK     int

	// Местоположение побочного максимума
	SideMaxValue float64
	SideMaxI     int
	SideMaxJ     int
	SideMaxK     int

	// Тип поля (объем, плоскости xy, yz, xz, оси x, y, z)
	Type string
}

func (field *Field) SetNodesNumber(nX, nY, nZ int) {
	field.Nx = nX
	field.Ny = nY
	field.Nz = nZ
}

func (field *Field) SetNodesSteps(dX, dY, dZ float64) {
	field.Dx = dX
	field.Dy = dY
	field.Dz = dZ
}

func (field *Field) SetGridBottomBorder(minX, minY, minZ float64) {
	field.MinX = minX
	field.MinY = minY
	field.MinZ = minZ
}

//This function prepares the Grid
func (field *Field) PrepareGrid() {
	field.X = make([]float64, field.Nx)
	field.Y = make([]float64, field.Ny)
	field.Z = make([]float64, field.Nz)

	// Так как вещественные числа не сохраняются точно в бинарном виде, то погрешность расчётов растёт при математических операциях.
	// В данном месте программы (задание сетки расчётов) наиболее важна точность координат сетки.
	// Для избежания ошибок вещественные числа умножаются на множитель FLOAT_ACCURACY_MLTPL = ,например, 1000000, преобразуются к целым,
	// производятся математические операции, а затем результат преобразуется обратно в вещественное число и делится на множитель FLOAT_ACCURACY_MLTPL.
	// Таким образом, при выборе множителя = 1000000 можно работать с точностью координат сетки до 1/1000000 метра, то есть до микрометра.

	for i := 0; i < field.Nx; i++ {
		field.X[i] = float64(int(field.MinX*FLOAT_ACCURACY_MLTPL)+i*int(field.Dx*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	for i := 0; i < field.Ny; i++ {
		field.Y[i] = float64(int(field.MinY*FLOAT_ACCURACY_MLTPL)+i*int(field.Dy*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	for i := 0; i < field.Nz; i++ {
		field.Z[i] = float64(int(field.MinZ*FLOAT_ACCURACY_MLTPL)+i*int(field.Dz*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	field.Value = make([]complex128, field.Nx*field.Ny*field.Nz)
	log.Println("Успешно создан массив точек поля размером", field.Nx*field.Ny*field.Nz)
}

//This function gets the indexes of 3D field and outputs the corresponding value at field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *Field) Get(i, j, k int) (v complex128) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	v = field.Value[ind]
	return
}

//This function gets the indexes of 3D field and puts the incoming value into field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *Field) Put(v complex128, i, j, k int) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	field.Value[ind] = v
}

//This function gets the indexes of 3D field and outputs the corresponding absolute value at field's point.
func (field *Field) Abs(i, j, k int) float64 {
	return cmplx.Abs(field.Get(i, j, k))
}

//This function gets the indexes of 3D field and outputs the corresponding phase value at field's point.
func (field *Field) Phase(i, j, k int) float64 {
	return cmplx.Phase(field.Get(i, j, k))
}

//This function finds maximum value of field and its indexes. Indexes fills Field struct
func (field *Field) MaxValue() (max float64) {
	for i := 0; i < field.Nx; i++ {
		for j := 0; j < field.Ny; j++ {
			for k := 0; k < field.Nz; k++ {
				if field.Abs(i, j, k) >= max {
					max = field.Abs(i, j, k)

				}
			}
		}
	}
	field.MainMaxValue = max
	return
}

func (field *Field) IndexesOfMaxValue() (I, J, K int) {
	max := 0.0
	for i := 0; i < field.Nx; i++ {
		for j := 0; j < field.Ny; j++ {
			for k := 0; k < field.Nz; k++ {
				if field.Abs(i, j, k) >= max {
					max = field.Abs(i, j, k)
					I = i
					J = j
					K = k
				}
			}
		}
	}
	return
}
