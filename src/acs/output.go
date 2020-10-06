// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var OutDir string

func AddDirToOutDir(oldDir string, newDirInt int) {
	newDir := strconv.Itoa(newDirInt)
	OutDir = filepath.Join(oldDir, newDir)
	error := os.Mkdir(OutDir, os.ModeDir)
	if error != nil {
		fmt.Printf("An error occurred on creating output folder %s\n", error)
		log.Printf("An error occurred on creating output folder %s\n", error)
		return
	}
	log.Println("Успешно создана папка для вывода", OutDir)
	return
}

//Generates output dir for current calculations titled with current date and time
func GenerateOutDirAndRedirectLog(workdir string) {
	workdir, _ = filepath.Abs(workdir)
	OutDir = filepath.Join(workdir, time.Now().Format("2006-01-02_15-04-05"))
	error := os.Mkdir(OutDir, os.ModeDir)
	if error != nil {
		fmt.Printf("An error occurred on creating output folder %s\n", error)
		log.Printf("An error occurred on creating output folder %s\n", error)
		return
	}
	redirectLog()
	log.Println("Успешно создана папка для вывода", OutDir)
}

func SetOutDirAndRedirectLog(workdir string) {
	OutDir, _ = filepath.Abs(workdir)
	redirectLog()
	log.Println("Задана папка для вывода", OutDir)
}

func redirectLog() {
	file := createTxtFileInOutDir("log")
	log.SetOutput(file)
	fmt.Println("Logging into", file.Name())
}

func PrintStatus(percents int, startTime time.Time) {
	fmt.Printf("Calculated: %d%%. ", percents)
	log.Printf("Рассчитано: %d%%\n", percents)
	if percents != 0 {
		d := time.Since(startTime)
		fmt.Println("Remaining:", time.Duration(100*int64(d)/int64(percents)-int64(d)))
	} else {
		fmt.Println("Can't define remaining time")
	}
}

func createFile(path string) (file *os.File) {
	file, erOpen := os.Create(path)
	if erOpen != nil {
		fmt.Printf("An error occurred on creating the files for output\n")
		log.Printf("An error occurred on creating the files for output\n")
	}
	return
}

func createGobFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".gob")
	return createFile(path)
}

func createBinFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".bin")
	return createFile(path)
}

func createTxtFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".txt")
	return createFile(path)
}

//This function saves (dumps) Field into file field.gob
func (field *Field) DumpField() {
	file := createGobFileInOutDir("field")
	defer file.Close()

	enc := gob.NewEncoder(file)
	erEnc := enc.Encode(field)
	if erEnc != nil {
		fmt.Printf("Error in dumping of Field %s\n", erEnc)
		log.Printf("Error in dumping of Field %s\n", erEnc)
		return
	}
	fmt.Println("Field dumping has acomplished")
	log.Println("Успешно закончено сохранение поля в", file.Name())
}

//This function restores Field from defined GOB-file
func (field *Field) RestoreFieldFromFile(filepath string) {
	file, erOpen := os.Open(filepath)
	if erOpen != nil {
		fmt.Printf("An error occurred on opening the file %s\n", erOpen)
		log.Printf("An error occurred on opening the file %s\n", erOpen)
		return
	}
	defer file.Close()

	erDec := gob.NewDecoder(file).Decode(&field)
	if erDec != nil {
		fmt.Printf("Error in restoring Field from file %s\n", erDec)
		return
	}

	fmt.Println("Restoring of field has finished")
	log.Println("Восстановление поля из файла закончено")
}

//This function prints planes XY, YZ, XZ of the Field into binary files (LittleEndian)
func (field *Field) PrintAbsFieldTxtPavel() {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	switch field.Type {
	case "yz":
		field.PrintAbsFieldYZtxtPavel(0)
	case "xz":
		field.PrintAbsFieldXZtxtPavel(0)
	case "xy":
		field.PrintAbsFieldXYtxtPavel(0)
	case "z":

	case "y":

	case "x":

	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

//This function prints planes XY, YZ, XZ of the Field into txt files with Pavel Rosnyckiy's template
func (field *Field) PrintAbsField2DtxtPavel(pointOnX int, pointOnY int, pointOnZ int) {
	fmt.Println("Text file Field output has started")
	log.Println("Начат вывод давления в текстовые файлы")

	field.PrintAbsFieldXYtxtPavel(pointOnZ)
	field.PrintAbsFieldYZtxtPavel(pointOnX)
	field.PrintAbsFieldXZtxtPavel(pointOnY)

	fmt.Println("Text file Field output has finished")
	log.Println("Вывод давления в текстовые файлы завершен")
}

func (field *Field) PrintAbsFieldYZtxtPavel(pointOnX int) {
	fmt.Println("Text file Field output has started")
	log.Println("Начат вывод давления в текстовые файлы")

	file := createTxtFileInOutDir("AbsField_YZ")
	defer file.Close()

	writeSizesTxtPavel(file, field.Y[0], field.Dy, field.Y[field.Ny-1], field.Z[0], field.Dz, field.Z[field.Nz-1], field.MaxValue())

	i := pointOnX
	for j, yj := range field.Y {
		for k, zk := range field.Z {
			writeValuesTxt(file, yj, zk, field.Abs(i, j, k))
		}
	}

	fmt.Println("Text file Field output has finished")
	log.Println("Вывод давления в текстовые файлы завершен")
}

func (field *Field) PrintAbsFieldXZtxtPavel(pointOnY int) {
	fmt.Println("Text file Field output has started")
	log.Println("Начат вывод давления в текстовые файлы")

	file := createTxtFileInOutDir("AbsField_XZ")
	defer file.Close()

	writeSizesTxtPavel(file, field.X[0], field.Dx, field.X[field.Nx-1], field.Z[0], field.Dz, field.Z[field.Nz-1], field.MaxValue())

	j := pointOnY
	for i, xi := range field.X {
		for k, zk := range field.Z {
			writeValuesTxt(file, xi, zk, field.Abs(i, j, k))
		}
	}
	fmt.Println("Text file Field output has finished")
	log.Println("Вывод давления в текстовые файлы завершен")
}

func (field *Field) PrintAbsFieldXYtxtPavel(pointOnZ int) {
	fmt.Println("Text file Field output has started")
	log.Println("Начат вывод давления в текстовые файлы")

	file := createTxtFileInOutDir("AbsField_XY")
	defer file.Close()

	writeSizesTxtPavel(file, field.X[0], field.Dx, field.X[field.Nx-1], field.Y[0], field.Dy, field.Y[field.Ny-1], field.MaxValue())

	k := pointOnZ
	for i, xi := range field.X {
		for j, yj := range field.Y {
			writeValuesTxt(file, xi, yj, field.Abs(i, j, k))
		}
	}

	fmt.Println("Text file Field output has finished")
	log.Println("Вывод давления в текстовые файлы завершен")
}

//This function prints planes XY, YZ, XZ of the Field into binary files (LittleEndian)
func (field *Field) PrintAbsFieldBinary(normalize bool) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	switch field.Type {
	case "yz":
		field.PrintAbsFieldYZBinary(normalize, 0)
	case "xz":
		field.PrintAbsFieldXZBinary(normalize, 0)
	case "xy":
		field.PrintAbsFieldXYBinary(normalize, 0)
	case "z":
		field.PrintAbsFieldZBinary(normalize, 0, 0)
	case "y":
		field.PrintAbsFieldYBinary(normalize, 0, 0)
	case "x":
		field.PrintAbsFieldXBinary(normalize, 0, 0)
	case "volume":
		field.PrintAbsField2DBinary(normalize, field.Nx/2, field.Ny/2, field.Nz/2)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

//This function prints planes XY, YZ, XZ of the Field into binary files (LittleEndian)
func (field *Field) PrintAbsField2DBinary(normalize bool, pointOnX int, pointOnY int, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	field.PrintAbsFieldXZBinary(normalize, pointOnY)
	field.PrintAbsFieldYZBinary(normalize, pointOnX)
	field.PrintAbsFieldXYBinary(normalize, pointOnZ)

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldYZBinary(normalize bool, pointOnX int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_YZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Ny, field.Nz)

	i := pointOnX
	for j, yj := range field.Y {
		for k, zk := range field.Z {
			writeValuesBin(file, yj, zk, field.Abs(i, j, k)*norm)
		}
	}

	writeFieldInfoBin(file, field)

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldXZBinary(normalize bool, pointOnY int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_XZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Nz)

	j := pointOnY
	for i, xi := range field.X {
		for k, zk := range field.Z {
			writeValuesBin(file, xi, zk, field.Abs(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldXYBinary(normalize bool, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_XY")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Ny)

	k := pointOnZ
	for i, xi := range field.X {
		for j, yj := range field.Y {
			writeValuesBin(file, xi, yj, field.Abs(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldZBinary(normalize bool, pointOnX, pointOnY int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_Z")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nz)

	i := pointOnX
	j := pointOnY
	for k, zk := range field.Z {
		writeValuesBin(file, zk, field.Abs(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldYBinary(normalize bool, pointOnX, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_Y")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Ny)

	i := pointOnX
	k := pointOnZ
	for j, yj := range field.Y {
		writeValuesBin(file, yj, field.Abs(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintAbsFieldXBinary(normalize bool, pointOnY, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("AbsField_X")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx)

	j := pointOnY
	k := pointOnZ
	for i, xi := range field.X {
		writeValuesBin(file, xi, field.Abs(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func writeFieldInfoBin(file *os.File, field *Field) {
	binary.Write(file, binary.LittleEndian, int64(field.FocusI))
	binary.Write(file, binary.LittleEndian, int64(field.FocusJ))
	binary.Write(file, binary.LittleEndian, int64(field.FocusK))

	binary.Write(file, binary.LittleEndian, int64(field.MainMaxI))
	binary.Write(file, binary.LittleEndian, int64(field.MainMaxJ))
	binary.Write(file, binary.LittleEndian, int64(field.MainMaxK))

	binary.Write(file, binary.LittleEndian, int64(field.SideMaxI))
	binary.Write(file, binary.LittleEndian, int64(field.SideMaxJ))
	binary.Write(file, binary.LittleEndian, int64(field.SideMaxK))
}

func writeSizesBin(file *os.File, values ...int) {
	for _, v := range values {
		binary.Write(file, binary.LittleEndian, int64(v))
	}
}

func writeValuesBin(file *os.File, values ...float64) {
	for _, v := range values {
		binary.Write(file, binary.LittleEndian, v)
	}
}

func writeSizesTxtPavel(file *os.File, values ...interface{}) {
	for _, v := range values {
		fmt.Fprintf(file, "%f\r\n", v)
	}
}

func writeValuesTxt(file *os.File, val1, val2, val3 float64) {
	fmt.Fprintf(file, "%f %f %f\r\n", val1, val2, val3)
}

//Функция создает файл, в который будет записываться информация в процессе сканирования фокуса и поиска бокового максимума
func CreateTxtForSideMaxInfo() *os.File {
	file := createTxtFileInOutDir("side_maxes")
	fmt.Fprintf(file, "Focus_X\tFocus_Y\tFocus_Z\tSideMax_X\tSideMax_Y\tSideMax_Z\tAmpInFocus\tAmpInSideMax\tDistaToScanPlane\tRatioMainSide\r\n")
	return file
}

//Функция записывает инфу в файл
func AddSideMaxInfo(file *os.File, x_f, y_f, z_f, x_s, y_s, z_s, amp_f, amp_s, dist, ratio float64) {
	fmt.Fprintf(file, "%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\r\n", x_f, y_f, z_f, x_s, y_s, z_s, amp_f, amp_s, dist, ratio)
}

func AddSideMaxInfoMATLAB(file *os.File, i_f, j_f, k_f, i_s, j_s, k_s int, amp_f, amp_s, ratio float64) {
	fmt.Fprintf(file, "%d\t%d\t%d\t%d\t%d\t%d\t%f\t%f\t%f\r\n", i_f, j_f, k_f, i_s, j_s, k_s, amp_f, amp_s, ratio)
}
