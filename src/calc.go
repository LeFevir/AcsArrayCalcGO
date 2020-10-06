// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package main

import (
	"./acs"
	"flag"
	"os"
	"time"
)

const (
	DEF_ARRAY_FILE        = "array.txt"
	DEF_ELEMENTS_FILE     = "array_elements.txt"
	DEF_ELEMENTS_OFF_FILE = "array_elements_off.txt"
	DEF_MEDIUM_FILE       = "medium.txt"
	DEF_GRID_CALC_FILE    = "grid_calc.txt"
	DEF_GRID_SCAN_FILE    = "grid_scan.txt"
	DEF_FOCUS_FILE        = "focus.txt"
)

var (
	procNum                                                                 int
	array_file, elements_file, elements_off_file, medium_file, focus_file   string
	grid_calc_file, grid_scan_file, work_folder, field_file, mult_grid_file string
	mode                                                                    string
	gendir, genbin, gengob, offel                                           bool
	oneElementCalc, normalize                                               bool
)

func main() {
	parseFlags()

	if gendir {
		acs.GenerateOutDirAndRedirectLog(work_folder)
	} else {
		acs.SetOutDirAndRedirectLog(work_folder)
	}

	array := acs.NewArrayFromFile(array_file)
	if oneElementCalc {
		array.AddElement(0, 0.0, 0.0)
	} else {
		array.AddElementsFromFile(elements_file)
	}
	if offel {
		array.RemoveElementsByIDs(elements_off_file)
	}
	medium := acs.NewMediumFromFile(medium_file)

	calcField := acs.NewGridFromFile(grid_calc_file)

	switch mode {

	case "calc":
		array.SetFocusFromFile(focus_file)
		calcField.AnalytCalcPressureField(procNum, array, medium, true)
		normalize = false

	case "scan-main":
		calcField.AnalytScanFocusForMainMaxes(array, medium)
		normalize = true

	case "scan-side":
		scanField := acs.NewGridFromFile(grid_scan_file)
		calcField.AnalytScanFocusForSideMaxes(procNum, scanField, array, medium)
		normalize = false

	case "multY":
		multField := acs.NewGridFromFile(mult_grid_file)
		outDir := acs.OutDir

		infoFile := acs.CreateTxtForSideMaxInfo()
		defer infoFile.Close()

		x_f := multField.MinX
		z_f := multField.MinZ

		count := multField.Ny
		startTime := time.Now()

		// Счетчик, чтобы правильно папки создавались
		counter := 0
		for j := 0; j < multField.Ny; j++ {
			y_f := multField.Y[j]

			counter += 1

			acs.AddDirToOutDir(outDir, counter)
			array.SetFocus(x_f, y_f, z_f)
			calcField.AnalytCalcPressureField(procNum, array, medium, true)
			calcField.MainSideMaxRationAndDistanceNEW(infoFile, x_f, y_f, z_f)
			calcField.PrintAbsFieldBinary(false)

			percents := 100 * (j) / count
			acs.PrintStatus(percents, startTime)
		}

	case "multZ":
		multField := acs.NewGridFromFile(mult_grid_file)
		outDir := acs.OutDir

		infoFile := acs.CreateTxtForSideMaxInfo()
		defer infoFile.Close()

		x_f := multField.MinX
		y_f := multField.MinY

		count := multField.Nz
		startTime := time.Now()

		// Счетчик, чтобы правильно папки создавались
		counter := 0

		for k := 0; k < multField.Nz; k++ {
			z_f := multField.Z[k]
			counter += 1

			acs.AddDirToOutDir(outDir, counter)
			array.SetFocus(x_f, y_f, z_f)
			calcField.AnalytCalcPressureField(procNum, array, medium, true)
			calcField.MainSideMaxRationAndDistanceNEW(infoFile, x_f, y_f, z_f)
			calcField.PrintAbsFieldBinary(false)

			percents := 100 * (k) / count
			acs.PrintStatus(percents, startTime)

		}
		os.Exit(0)

	case "multYZ":
		multField := acs.NewGridFromFile(mult_grid_file)
		outDir := acs.OutDir

		infoFile := acs.CreateTxtForSideMaxInfo()
		defer infoFile.Close()

		x_f := multField.MinX

		count := multField.Nz * multField.Ny * multField.Nx
		startTime := time.Now()

		// Странный цикл для красивого скольжения луча снизу вверх,
		// а потом сверху вниз, то есть змейкой.
		// Возможно можно переписать красивее

		// Счетчик, чтобы правильно папки создавались
		counter := 0

		for k := 0; k < multField.Nz; k++ {
			z_f := multField.Z[k]
			for j := 0; j < multField.Ny; j++ {
				y_f := multField.Y[j]

				counter += 1

				acs.AddDirToOutDir(outDir, counter)
				array.SetFocus(x_f, y_f, z_f)
				calcField.AnalytCalcPressureField(procNum, array, medium, true)
				calcField.MainSideMaxRationAndDistanceNEW(infoFile, x_f, y_f, z_f)
				calcField.PrintAbsFieldBinary(false)

				percents := 100 * (j + k*multField.Ny) / count
				acs.PrintStatus(percents, startTime)
			}

			k++
			z_f = multField.Z[k]
			for j := multField.Ny - 1; j >= 0; j-- {
				y_f := multField.Y[j]

				counter += 1
				acs.AddDirToOutDir(outDir, counter)
				array.SetFocus(x_f, y_f, z_f)
				calcField.AnalytCalcPressureField(procNum, array, medium, true)
				calcField.MainSideMaxRationAndDistanceNEW(infoFile, x_f, y_f, z_f)
				calcField.PrintAbsFieldBinary(false)

				percents := 100 * (j + k*multField.Ny) / count
				acs.PrintStatus(percents, startTime)
			}
		}
		os.Exit(0)

	case "exactSolutionZ":
		calcField.CalcExactSolutionOnZ(array, medium)
		normalize = false

	case "numCalc":
		calcField.NumericCalcPressureField(procNum, array, medium, 1000, true)
		normalize = false
	}

	if genbin {
		calcField.PrintAbsFieldBinary(normalize)
	}

	if gengob {
		calcField.DumpField()
	}

	calcField.PrintAbsFieldTxtPavel()
}

func parseFlags() {
	flag.IntVar(&procNum, "proc", 0, "number of threads to use. 0 for autodetection")
	flag.StringVar(&mode, "mode", "calc", "mode:\ncalc - calculate field with focus in point which has set in focus.txt\nscan-main - scan focus for main maxes amplitudes\nmultY - create series of fields with steering focus by coordinates from mult")
	flag.StringVar(&array_file, "array", DEF_ARRAY_FILE, "path to txt file with array params")
	flag.StringVar(&elements_file, "els", DEF_ELEMENTS_FILE, "path to txt file with array elements coords")
	flag.StringVar(&medium_file, "medium", DEF_MEDIUM_FILE, "path to txt file with medium params")
	flag.StringVar(&grid_calc_file, "gridcalc", DEF_GRID_CALC_FILE, "path to txt file with calculation grid params")
	flag.StringVar(&grid_scan_file, "gridscan", "", "optional path to txt file with scanning grid params")
	flag.StringVar(&focus_file, "focus", DEF_FOCUS_FILE, "optional path to txt file with focus coords")
	flag.StringVar(&mult_grid_file, "gridmult", "", "optional path to txt file with mult grid")
	flag.StringVar(&work_folder, "dir", "", "optional path to output folder")
	flag.BoolVar(&gendir, "gendir", false, "generate or not output folder with time stamp")
	flag.BoolVar(&genbin, "genbin", false, "generate or not binary field files")
	flag.BoolVar(&gengob, "gengob", false, "dump or not field gob file")
	flag.BoolVar(&oneElementCalc, "one", false, "calc field for one element")
	flag.BoolVar(&offel, "offel", false, "turn off elements or not")
	flag.StringVar(&field_file, "field", "", "optional path to field gob-file to restore the field without calculating")
	flag.StringVar(&elements_off_file, "off", DEF_ELEMENTS_OFF_FILE, "path to txt file with list of turning off elements")
	flag.Parse()
}
