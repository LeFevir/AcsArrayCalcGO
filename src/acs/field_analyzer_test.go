// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import (
	"testing"
)

func TestFindFieldPointByCoord(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_x.txt")
	i,_,_:=field.findFieldPointByCoord(-0.0242,0.0,0.0)
	if i!=1 {
		t.Log("Error in finding Field Point on X!", i, field.X[i])
		t.Fail()
	}
	
	field = NewGridFromFile("test_files\\grid_y.txt")
	_,j,_:=field.findFieldPointByCoord(0.0,-0.0242,0.0)
	if j!=1 {
		t.Log("Error in finding Field Point on X!", j, field.Y[j])
		t.Fail()
	}
	
	field = NewGridFromFile("test_files\\grid_z.txt")
	_,_,k:=field.findFieldPointByCoord(0.0,0.0,0.071)
	if j!=1 {
		t.Log("Error in finding Field Point on X!", k, field.Z[k])
		t.Fail()
	}
}
