// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

//The package ACS includes the algorithms and useful functions for calculating the acoustic field of
//multi-element array with circular elements.

import (
	"fmt"
	"log"
	"math"
)

//Acoustic array structure
type Array struct {
	Name            string
	Aperture        float64
	CurvatureRadius float64
	ElementRadius   float64
	Frequency       float64
	Elements        []Element
}

//Element of array structure.
//ID - identification of element (number).
//Center - coordinates of center of element
//Ro - distance from center of element to center of array
//Cos_Phi, Sin_Phi - COS and SIN of PHI - polar angle of element's center according to canter of array
//Cos_Gamma, Sin_Gamma - COS and SIN of GAMMA - azimuth angle of element's center according to canter of curvature of array
//PhaseShift - distance from the element's center to focus. Needs for electronic scanning.
type Element struct {
	ID         int
	Center     Point
	Cos_Phi    float64
	Sin_Phi    float64
	Cos_Gamma  float64
	Sin_Gamma  float64
	PhaseShift float64
}

func NewElementOfArray(array *Array, id int, coordX, coordY, coordZ float64) (element Element) {
	element.ID = id
	element.Center = NewPoint(coordX, coordY, coordZ)
	element.CalcElementParameters(array.CurvatureRadius)
	return
}

//This Function calculates additional parameters of the element for performance benefit
func (el *Element) CalcElementParameters(arrayCurvatureRadius float64) {
	ro := math.Hypot(el.Center.X, el.Center.Y)
	el.Cos_Phi = el.Center.Y / ro
	el.Sin_Phi = el.Center.X / ro

	x_gamma := ro
	y_gamma := arrayCurvatureRadius - el.Center.Z
	hyp_gamma := math.Hypot(x_gamma, y_gamma)
	el.Cos_Gamma = y_gamma / hyp_gamma
	el.Sin_Gamma = x_gamma / hyp_gamma
	return
}

//This Function outputs the number of Elements
func (array *Array) NumberOfElements() int {
	return len(array.Elements)
}

//This Function adds an Element of Acoustic Array with ID, X and Y coordinates
func (array *Array) AddElement(id int, coordX, coordY float64) {
	if (math.Hypot(coordX, coordY) > array.Aperture/2) || (math.Hypot(coordX, coordY) > array.CurvatureRadius) {
		fmt.Println("Wrong coordinates of element", id)
		log.Println("Неправильные координаты элемента", id)
		return
	}
	//Calculating Z coordinate corresponding to CurvatureRadius
	coordZ := array.CurvatureRadius - math.Sqrt(array.CurvatureRadius*array.CurvatureRadius-coordX*coordX-coordY*coordY)
	//Adding new element to the Array of Elements
	element := NewElementOfArray(array, id, coordX, coordY, coordZ)
	array.Elements = append(array.Elements, element)
}

//This Function removes an Element of Acoustic Array by its ID
func (array *Array) removeElementByID(id int) {
	if len(array.Elements) == 0 {
		return
	}
	for i := 0; i < len(array.Elements); i++ {
		if array.Elements[i].ID == id {
			array.Elements = append(array.Elements[:i], array.Elements[i+1:]...)
			return
		}
	}
	return
}

//This Function removes thr Elements of Acoustic Array by its IDS
func (array *Array) RemoveElementsByIDs(path string) {
	ids := ReadTurnOffElementsFromFile(path)
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		array.removeElementByID(id)
	}
	fmt.Println("Elements have been deleted. Now there are", array.NumberOfElements(), "elements")
	log.Println("Элементы были удалены. Теперь их", array.NumberOfElements())
	return
}

//This Function sets the Electronic focus into defined point with fX,fY,fZ coordinated
func (array *Array) SetFocus(fX, fY, fZ float64) {
	for i := 0; i < array.NumberOfElements(); i++ {
		element := array.Elements[i]
		array.Elements[i].PhaseShift = math.Sqrt(math.Pow((fX-element.Center.X), 2) + math.Pow((fY-element.Center.Y), 2) + math.Pow((fZ-element.Center.Z), 2))
	}
}
