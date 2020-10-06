// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"math/cmplx"
	"testing"
)

func TestCircTransducerFarField(t *testing.T) {
	// Test for zero value of theta
	transRadius := 3.3e-03
	waveNumber := 12.0
	theta := 0.0
	r := 100e-03
	x := circTransducerFarField(transRadius, waveNumber, theta, r)
	if cmplx.IsNaN(x) {
		t.Log("Wrong Far Field calc", x)
		t.Fail()
	}
}
