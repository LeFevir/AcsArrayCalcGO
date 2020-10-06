// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import ()

type triangle struct {
	V1, V2, V3 Point  //Vertexes of triangle
	Center     Point
	Area       float64
}

func NewTriangle(x1, y1, x2, y2, x3, y3 float64) (tri triangle) {
	tri.V1 = NewPoint(x1, y1, 0.0)
	tri.V2 = NewPoint(x2, y2, 0.0)
	tri.V3 = NewPoint(x3, y3, 0.0)
	return
}
