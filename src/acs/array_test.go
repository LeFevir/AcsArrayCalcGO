// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import (
	"math"
	"testing"
)

func TestAddElement(t *testing.T) {
	array := new(Array)
	array.Aperture = 6.0
	array.CurvatureRadius = 3.0
	array.AddElement(1, 1.0, 2.0)

	if array.NumberOfElements() != 1 {
		t.Log("Wrong element adding!")
		t.Fail()
	} else {
		z := array.CurvatureRadius - math.Sqrt(array.CurvatureRadius*array.CurvatureRadius-1.0-4.0)
		if array.Elements[0].Center.Z != z {
			t.Log("Wrong z coordinate calc! Z =", z)
			t.Fail()
		}
	}
}

func TestRemoveElementByID(t *testing.T) {
	var array Array
	array.Aperture = 6.0
	array.CurvatureRadius = 3.0
	array.AddElement(0, 0.0, 0.0)
	array.AddElement(1, 1.0, 1.0)
	array.AddElement(2, 2.0, 2.0)

	array.RemoveElementByID(0)
	if array.NumberOfElements() != 2 {
		t.Log("Wrong removing!")
		t.Fail()
	} else if (array.Elements[0].ID != 1) || (array.Elements[1].ID != 2) {
		t.Log("Wrong removing!")
		t.Fail()
	}

	var array2 Array
	array2.Aperture = 6.0
	array2.CurvatureRadius = 3.0
	array2.AddElement(0, 0.0, 0.0)
	array2.AddElement(1, 1.0, 1.0)
	array2.AddElement(2, 2.0, 2.0)
	array2.RemoveElementByID(2)
	if array2.NumberOfElements() != 2 {
		t.Log("Wrong removing!")
		t.Fail()
	} else if (array2.Elements[0].ID != 0) || (array2.Elements[1].ID != 1) {
		t.Log("Wrong removing!")
		t.Fail()
	}
}
