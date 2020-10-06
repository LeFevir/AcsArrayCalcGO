// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"fmt"
	"log"
	"time"
)

//Function using scanning field to set focus in point of that field and calculate analytically
//pressure in this point.
func (scanField *Field) AnalytScanFocusForMainMaxes(array *Array, medium *Medium) {
	//TODO Сделать многопоточность. Если в лоб - фокус неправильно расставляется по потокам
	fmt.Println("Focus scanning for Main Maxes has started")
	log.Println("Начато сканирование фокуса для основных максимумов")

	var p complex128
	for k, z_f := range scanField.Z {
		for j, y_f := range scanField.Y {
			for i, x_f := range scanField.X {
				//Set phase focus into point of scanning field
				array.SetFocus(x_f, y_f, z_f)
				//Calculate the pressure in point of scanning field
				p = AnalytCalcElementsField(array, medium.SpeedOfSound, NewPoint(x_f, y_f, z_f))
				//Save amplitude of Main Max in scanning field
				scanField.Put(p, i, j, k)
			}
		}
	}
}

func (scanField *Field) AnalytScanFocusForSideMaxes(procNum int, calcField *Field, array *Array, medium *Medium) {

	fmt.Println("Focus scanning for Side Maxes has started")
	log.Println("Началось сканирование фокуса для боковых максимумов")

	infoFile := CreateTxtForSideMaxInfo()
	defer infoFile.Close()

	count := scanField.Nz * scanField.Ny * scanField.Nx
	startTime := time.Now()
	for k, z_f := range scanField.Z {
		for j, y_f := range scanField.Y {
			for i, x_f := range scanField.X {
				//Set phase focus into point of scanning field
				array.SetFocus(x_f, y_f, z_f)
				//Calculate the pressure in volume of calc field
				calcField.AnalytCalcPressureField(procNum, array, medium, true)
				ratio := calcField.MainSideMaxRationAndDistanceNEW(infoFile, x_f, y_f, z_f)
				//Save Ratio in scanning field
				scanField.Put(complex(ratio, 0.0), i, j, k)

				percents := 100 * (i + j*scanField.Nx + k*scanField.Ny*scanField.Nx) / count
				PrintStatus(percents, startTime)
			}
		}
	}

	fmt.Printf("Focus scanning for Side Maxes has finished. It took %s\n", time.Since(startTime))
	log.Printf("Сканирование фокуса для боковых максимумов закончено. Оно заняло %s\n", time.Since(startTime))

}
