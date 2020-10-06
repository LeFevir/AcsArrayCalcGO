// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func openFile(path string) *os.File {
	path, _ = filepath.Abs(path)
	file, erOpen := os.Open(path)
	if erOpen != nil {
		fmt.Printf("An error occurred on opening the file, %s\n", erOpen)
		log.Printf("An error occurred on opening the file, %s\n", erOpen)
		os.Exit(1)
	}
	return file
}

func parseParamFile(path string) (values []string) {
	file := openFile(path)
	defer file.Close()
	fileReader := skipByteOrderMarkIfNeed(bufio.NewReader(file))
	for {
		s, erRead := fileReader.ReadString('\n')
		s2 := strings.Split(s, ":")
		if len(s2) == 0 {
			values = append(values, strings.TrimSpace(s))
		} else {
			values = append(values, strings.TrimSpace(s2[len(s2)-1]))
		}
		if erRead != nil {
			break
		}
	}
	return
}

func skipByteOrderMarkIfNeed(fileReader *bufio.Reader) *bufio.Reader {
	b1, _ := fileReader.ReadByte()
	b2, _ := fileReader.ReadByte()
	b3, _ := fileReader.ReadByte()
	if (b1 != 0xEF) && (b2 != 0xBB) && (b3 != 0xBF) {
		fileReader.UnreadByte()
		fileReader.UnreadByte()
		fileReader.UnreadByte()
	}
	return fileReader
}

func NewArrayFromFile(path string) (array *Array) {
	values := parseParamFile(path)
	if len(values) < 5 {
		fmt.Println("Wrong Array file!")
		log.Println("Неправильные файл решетки!")
		os.Exit(1)
	}

	array = new(Array)
	array.Name = values[0]
	array.Aperture, _ = strconv.ParseFloat(values[1], 64)
	array.CurvatureRadius, _ = strconv.ParseFloat(values[2], 64)
	array.ElementRadius, _ = strconv.ParseFloat(values[3], 64)
	array.Frequency, _ = strconv.ParseFloat(values[4], 64)

	fmt.Println("Array parameters has read for", array.Name, "array")
	log.Println("Успешно считан файл решетки", array.Name)
	return
}

//This Function reads Elements file and fills the Array Elements structure
func (array *Array) AddElementsFromFile(path string) {
	file := openFile(path)
	defer file.Close()
	fileReader := bufio.NewReader(file)
	skipByteOrderMarkIfNeed(fileReader)

	var x, y float64

	for i := 0; ; i++ {
		_, err := fmt.Fscanln(fileReader, &x, &y)
		if err != nil {
			break
		}
		array.AddElement(i, x, y)
	}
	fmt.Println("Elements of Array have read, total:", array.NumberOfElements())
	log.Println("Успешно считаны координаты элементов, всего элементов", array.NumberOfElements())
}

//This Function reads Medium Info file and creates new Medium structure
func NewMediumFromFile(path string) (medium *Medium) {
	values := parseParamFile(path)
	if len(values) < 3 {
		fmt.Println("Wrong medium file!")
		log.Println("Неправильные файл среды!")
		os.Exit(1)
	}
	medium = new(Medium)
	medium.Name = values[0]
	medium.Density, _ = strconv.ParseFloat(values[1], 64)
	medium.SpeedOfSound, _ = strconv.ParseFloat(values[2], 64)

	fmt.Println("Medium parameters have filled for", medium.Name)
	log.Println("Успешно считан файл с параметрами среды", medium.Name)
	return
}

func (array *Array) SetFocusFromFile(path string) {
	values := parseParamFile(path)
	if len(values) != 3 {
		fmt.Println("Wrong focus file!")
		log.Println("Неправильный файл с фокусом!")
		os.Exit(1)
	}
	x, _ := strconv.ParseFloat(values[0], 64)
	y, _ := strconv.ParseFloat(values[1], 64)
	z, _ := strconv.ParseFloat(values[2], 64)
	array.SetFocus(x, y, z)
	fmt.Println("Focus has set into", x, y, z)
	log.Println("Фокус установлен в", x, y, z)
}

func ReadTurnOffElementsFromFile(path string) (ids []int) {
	values := parseParamFile(path)
	for _, v := range values {
		id, _ := strconv.Atoi(v)
		ids = append(ids, id)
	}
	return
}

func NewGridFromFile(path string) (field *Field) {
	values := parseParamFile(path)
	switch values[0] {
	case "volume":
		field = NewVolGridFromParams(values[1:])
		field.Type = "volume"
	case "yz":
		field = NewYZGridFromParams(values[1:])
		field.Type = "yz"
	case "xz":
		field = NewXZGridFromParams(values[1:])
		field.Type = "xz"
	case "xy":
		field = NewXYGridFromParams(values[1:])
		field.Type = "xy"
	case "z":
		field = NewZGridFromParams(values[1:])
		field.Type = "z"
	case "y":
		field = NewYGridFromParams(values[1:])
		field.Type = "y"
	case "x":
		field = NewXGridFromParams(values[1:])
		field.Type = "x"
	default:
		fmt.Println("Wrong grid file! No Grid Type line!")
		log.Println("Неправильный файл сетки! Нет строки типа сетки!")
		os.Exit(1)
	}
	fmt.Println("Grid parameters has filled successful")
	log.Println("Успешно считаны параметры сетки")
	return
}

func NewVolGridFromParams(values []string) (field *Field) {
	if len(values) != 9 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Nx, _ = strconv.Atoi(values[0])
	field.Ny, _ = strconv.Atoi(values[1])
	field.Nz, _ = strconv.Atoi(values[2])
	field.MinX, _ = strconv.ParseFloat(values[3], 64)
	field.MinY, _ = strconv.ParseFloat(values[4], 64)
	field.MinZ, _ = strconv.ParseFloat(values[5], 64)
	field.Dx, _ = strconv.ParseFloat(values[6], 64)
	field.Dy, _ = strconv.ParseFloat(values[7], 64)
	field.Dz, _ = strconv.ParseFloat(values[8], 64)

	field.PrepareGrid()
	return
}

func NewYZGridFromParams(values []string) (field *Field) {
	if len(values) != 7 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Nx = 1
	field.Dx = 1.0
	field.MinX, _ = strconv.ParseFloat(values[0], 64)
	field.Ny, _ = strconv.Atoi(values[1])
	field.Nz, _ = strconv.Atoi(values[2])
	field.MinY, _ = strconv.ParseFloat(values[3], 64)
	field.MinZ, _ = strconv.ParseFloat(values[4], 64)
	field.Dy, _ = strconv.ParseFloat(values[5], 64)
	field.Dz, _ = strconv.ParseFloat(values[6], 64)

	field.PrepareGrid()
	return
}

func NewXZGridFromParams(values []string) (field *Field) {
	if len(values) != 7 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Ny = 1
	field.Dy = 1.0
	field.MinY, _ = strconv.ParseFloat(values[0], 64)
	field.Nx, _ = strconv.Atoi(values[1])
	field.Nz, _ = strconv.Atoi(values[2])
	field.MinX, _ = strconv.ParseFloat(values[3], 64)
	field.MinZ, _ = strconv.ParseFloat(values[4], 64)
	field.Dx, _ = strconv.ParseFloat(values[5], 64)
	field.Dz, _ = strconv.ParseFloat(values[6], 64)

	field.PrepareGrid()
	return
}

func NewXYGridFromParams(values []string) (field *Field) {
	if len(values) != 7 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Nz = 1
	field.Dz = 1.0
	field.MinZ, _ = strconv.ParseFloat(values[0], 64)
	field.Nx, _ = strconv.Atoi(values[1])
	field.Ny, _ = strconv.Atoi(values[2])
	field.MinX, _ = strconv.ParseFloat(values[3], 64)
	field.MinY, _ = strconv.ParseFloat(values[4], 64)
	field.Dx, _ = strconv.ParseFloat(values[5], 64)
	field.Dy, _ = strconv.ParseFloat(values[6], 64)

	field.PrepareGrid()
	return
}

func NewZGridFromParams(values []string) (field *Field) {
	if len(values) != 5 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Nx = 1
	field.Dx = 1.0
	field.Ny = 1
	field.Dy = 1.0
	field.MinX, _ = strconv.ParseFloat(values[0], 64)
	field.MinY, _ = strconv.ParseFloat(values[1], 64)
	field.Nz, _ = strconv.Atoi(values[2])
	field.MinZ, _ = strconv.ParseFloat(values[3], 64)
	field.Dz, _ = strconv.ParseFloat(values[4], 64)

	field.PrepareGrid()
	return
}

func NewYGridFromParams(values []string) (field *Field) {
	if len(values) != 5 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Nx = 1
	field.Dx = 1.0
	field.Nz = 1
	field.Dz = 1.0
	field.MinX, _ = strconv.ParseFloat(values[0], 64)
	field.MinZ, _ = strconv.ParseFloat(values[1], 64)
	field.Ny, _ = strconv.Atoi(values[2])
	field.MinY, _ = strconv.ParseFloat(values[3], 64)
	field.Dy, _ = strconv.ParseFloat(values[4], 64)

	field.PrepareGrid()
	return
}

func NewXGridFromParams(values []string) (field *Field) {
	if len(values) != 5 {
		fmt.Println("Wrong grid file!")
		log.Println("Неправильный файл с сеткой расчетов!")
		os.Exit(1)
	}
	field = new(Field)
	field.Ny = 1
	field.Dy = 1.0
	field.Nz = 1
	field.Dz = 1.0
	field.MinY, _ = strconv.ParseFloat(values[0], 64)
	field.MinZ, _ = strconv.ParseFloat(values[1], 64)
	field.Nx, _ = strconv.Atoi(values[2])
	field.MinX, _ = strconv.ParseFloat(values[3], 64)
	field.Dx, _ = strconv.ParseFloat(values[4], 64)

	field.PrepareGrid()
	return
}
