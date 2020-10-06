// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"time"
)

func (field *Field) CalcExactSolutionOnZ(array *Array, medium *Medium) {

	fmt.Println("Exact Solution on Z calculation has started")
	log.Println("Расчет точного решения по оси Z")

	startTime := time.Now()

	for k, zk := range field.Z {
		p := exactSolutionOnZ(array, medium.SpeedOfSound, zk)
		field.Put(p, 0, 0, k)
	}

	fmt.Printf("Exact Solution on Z calculation has finished. It took %s\n", time.Since(startTime))
	log.Printf("Расчет точного решения по оси Z закончен. Он занял %s\n", time.Since(startTime))
}

//Calculates field in point from all active Elements of array
func exactSolutionOnZ(array *Array, speedOfSound float64, z float64) (p complex128) {
	k := 2 * math.Pi * array.Frequency / speedOfSound
	s1 := (k / 2) * (math.Sqrt(array.ElementRadius*array.ElementRadius+z*z) + z)
	s2 := (k / 2) * (math.Sqrt(array.ElementRadius*array.ElementRadius+z*z) - z)
	p = (-2i) * cmplx.Exp(complex(0.0, s1)) * cmplx.Sin(complex(s2, 0.0))
	return
}
