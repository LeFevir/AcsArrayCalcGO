// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"fmt"
	"math"
	"os"
)

const (
	MAIN_MAX_SPOT_DX = 5.0e-03
	MAIN_MAX_SPOT_DY = 5.0e-03
	MAIN_MAX_SPOT_DZ = 15.0e-03
)

//Функция должна найти точку в вызваемом поле, соответствующую координатам из аргументов
func (field *Field) findFieldPointByCoord(x, y, z float64) (diffI, diffJ, diffK int) {

	diff := 1.0
	for i := 0; i < field.Nx; i++ {
		if math.Abs(x-field.X[i]) <= diff {
			diff = math.Abs(x - field.X[i])
			diffI = i
			if diff == 0.0 { //оптимизация
				break
			}
		}
	}

	diff = 1.0
	for j := 0; j < field.Ny; j++ {
		if math.Abs(y-field.Y[j]) <= diff {
			diff = math.Abs(y - field.Y[j])
			diffJ = j
			if diff == 0.0 {
				break
			}
		}
	}

	diff = 1.0
	for k := 0; k < field.Nz; k++ {
		if math.Abs(z-field.Z[k]) <= diff {
			diff = math.Abs(z - field.Z[k])
			diffK = k
			if diff == 0.0 {
				break
			}
		}
	}
	return
}

//TODO переписать!!!
func (field *Field) MainSideMaxRationAndDistance(infoFile *os.File, thresholdRatio float64, x_f, y_f, z_f float64) (float64, float64) {
	//	threshold := thresholdRatio * field.MaxValue()
	//	maxi, maxj, maxk := field.IndexesOfMaxValue()

	maxi, maxj, maxk := field.findFieldPointByCoord(x_f, y_f, z_f)
	threshold := thresholdRatio * field.Abs(maxi, maxj, maxk)

	//	threshold := thresholdRatio * field.MaxValue()
	//	maxi, maxj, maxk := mi, mj, mk

	i := maxi
	j := maxj
	k := maxk
	spotXMin := 0
	for i = maxi; i > 0; i-- {
		if field.Abs(i, j, k) < threshold {
			spotXMin = i
			break
		}
	}
	fmt.Println(spotXMin, field.X[spotXMin])
	spotXMax := field.Nx - 1
	for i = maxi; i < field.Nx; i++ {
		if field.Abs(i, j, k) < threshold {
			spotXMax = i
			break
		}
	}
	fmt.Println(spotXMax, field.X[spotXMax])

	spotYMin := 0
	i = maxi
	j = maxj
	k = maxk
	for j = maxj; j > 0; j-- {
		if field.Abs(i, j, k) < threshold {
			spotYMin = j
			break
		}
	}
	fmt.Println(spotYMin, field.Y[spotYMin])
	spotYMax := field.Ny - 1
	for j = maxj; j < field.Ny; j++ {
		if field.Abs(i, j, k) < threshold {
			spotYMax = j
			break
		}
	}
	fmt.Println(spotYMax, field.Y[spotYMax])

	spotZMin := 0
	i = maxi
	j = maxj
	k = maxk
	for k = maxk; k > 0; k-- {
		if field.Abs(i, j, k) < threshold {
			spotZMin = k
			break
		}
	}
	fmt.Println(spotZMin, field.Z[spotZMin])
	spotZMax := field.Nz - 1
	for k = maxk; k < field.Nz; k++ {
		if field.Abs(i, j, k) < threshold {
			spotZMax = k
			break
		}
	}
	fmt.Println(spotZMax, field.Z[spotZMax])

	field.SideMaxValue = 0.0
	for i = 0; i < field.Nx; i++ {
		for j = 0; j < field.Ny; j++ {
			for k = 0; k < field.Nz; k++ {
				if !((i >= spotXMin) && (i <= spotXMax) && (j >= spotYMin) && (j <= spotYMax) && (k >= spotZMin) && (k <= spotZMax)) {
					if field.Abs(i, j, k) > field.SideMaxValue {
						field.SideMaxValue = field.Abs(i, j, k)
						field.SideMaxI = i
						field.SideMaxJ = j
						field.SideMaxK = k
					}
				}
			}
		}
	}

	fmt.Println(field.X[field.SideMaxI], field.Y[field.SideMaxJ], field.Z[field.SideMaxK])
	fmt.Println(field.SideMaxValue, field.Abs(maxi, maxj, maxk), field.SideMaxValue/field.Abs(maxi, maxj, maxk))

	dist := calcDistToScanPlane(x_f, y_f, field.X[field.SideMaxI], field.Y[field.SideMaxJ])
	ratio := field.SideMaxValue / field.Abs(maxi, maxj, maxk)
	AddSideMaxInfo(infoFile, x_f, y_f, z_f, field.X[field.SideMaxI], field.Y[field.SideMaxJ], field.Z[field.SideMaxK], field.Abs(maxi, maxj, maxk), field.SideMaxValue, dist, ratio)

	return ratio, dist
}

//Calculating the distance from scan plane (with Z axis and focal point) to side max
func calcDistToScanPlane(x_focus, y_focus, x_side_max, y_side_max float64) float64 {

	if x_focus == 0.0 {
		return math.Abs(x_side_max)
	}

	return math.Abs((y_focus/x_focus)*x_side_max-y_side_max) / math.Sqrt(1+(y_focus/x_focus)*(y_focus/x_focus))
}

//TODO переписать!!!
func (field *Field) MainSideMaxRationAndDistanceNEW(infoFile *os.File, x_f, y_f, z_f float64) float64 {

	// Записываем координаты фокуса
	field.FocusI, field.FocusJ, field.FocusK = field.findFieldPointByCoord(x_f, y_f, z_f)

	maxi, maxj, maxk := field.IndexesOfMaxValue() //Для таких случаев, когда фокус портится и происходит раздвоение

	field.MainMaxI = maxi
	field.MainMaxJ = maxj
	field.MainMaxK = maxk

	spotXMin, _, _ := field.findFieldPointByCoord(x_f-MAIN_MAX_SPOT_DX, y_f, z_f)
	// fmt.Println("spotXMin =", spotXMin, field.X[spotXMin])

	spotXMax, _, _ := field.findFieldPointByCoord(x_f+MAIN_MAX_SPOT_DX, y_f, z_f)
	// fmt.Println("spotXMax =", spotXMax, field.X[spotXMax])

	_, spotYMin, _ := field.findFieldPointByCoord(x_f, y_f-MAIN_MAX_SPOT_DY, z_f)
	// fmt.Println("spotYMin =", spotYMin, field.Y[spotYMin])

	_, spotYMax, _ := field.findFieldPointByCoord(x_f, y_f+MAIN_MAX_SPOT_DY, z_f)
	// fmt.Println("spotYMax =", spotYMax, field.Y[spotYMax])

	_, _, spotZMin := field.findFieldPointByCoord(x_f, y_f, z_f-MAIN_MAX_SPOT_DZ)
	// fmt.Println("spotZMin =", spotZMin, field.Z[spotZMin])

	_, _, spotZMax := field.findFieldPointByCoord(x_f, y_f, z_f+MAIN_MAX_SPOT_DZ)
	// fmt.Println("spotZMax =", spotZMax, field.Z[spotZMax])

	field.SideMaxValue = 0.0
	for i := 0; i < field.Nx; i++ {
		for j := 0; j < field.Ny; j++ {
			for k := 0; k < field.Nz; k++ {
				if !((i >= spotXMin) && (i <= spotXMax) && (j >= spotYMin) && (j <= spotYMax) && (k >= spotZMin) && (k <= spotZMax)) {
					if field.Abs(i, j, k) > field.SideMaxValue {
						field.SideMaxValue = field.Abs(i, j, k)
						field.SideMaxI = i
						field.SideMaxJ = j
						field.SideMaxK = k
					}
				}
			}
		}
	}

	//fmt.Println(field.X[field.SideMaxI], field.Y[field.SideMaxJ], field.Z[field.SideMaxK])
	//fmt.Println(field.SideMaxValue, field.Abs(maxi, maxj, maxk), field.SideMaxValue/field.Abs(maxi, maxj, maxk))

	dist := calcDistToScanPlane(x_f, y_f, field.X[field.SideMaxI], field.Y[field.SideMaxJ])
	ratio := field.SideMaxValue / field.Abs(maxi, maxj, maxk)
	AddSideMaxInfo(infoFile, x_f, y_f, z_f, field.X[field.SideMaxI], field.Y[field.SideMaxJ], field.Z[field.SideMaxK], field.Abs(maxi, maxj, maxk), field.SideMaxValue, dist, ratio)
	//AddSideMaxInfoMATLAB(infoFile, maxi, maxj, maxk, field.SideMaxI, field.SideMaxJ, field.SideMaxK, field.Abs(maxi, maxj, maxk), field.SideMaxValue, ratio)

	return ratio
}
