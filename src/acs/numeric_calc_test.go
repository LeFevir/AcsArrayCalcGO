// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import (
	"testing"
)

func TestSliceOnTrisProper(t *testing.T) {
	array := new(Array)
	array.Aperture = 4.0
	array.CurvatureRadius = 3.0
	array.ElementRadius = 1.0
	array.AddElement(0, 0.0, 0.0)
	tris := array.sliceOnTrisProper(8)

	if len(tris) != 8 {
		t.Log("Wrong number of tris after slicing", len(tris))
		t.Fail()
	}
}
