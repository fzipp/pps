// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "math"

// sign returns the sign, also called signum, of i: -1 for a negative number,
// 0 for the number zero, and +1 for a positive number.
func sign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return +1
	}
	return 0
}

// deg2rad converts the measurement of an angle from degrees to radians.
func deg2rad(deg float64) (rad float64) {
	return deg * (math.Pi / 180)
}

func wrap(n, max float64) float64 {
	if n > max {
		return n - max
	}
	if n < 0 {
		return max + n
	}
	return n
}

const epsilon = 1e-10

// nearEq compares two floating-point numbers for equality within an
// absolute difference tolerance of epsilon.
// This relation is not transitive, except for ε=0.
func nearEq(a, b, ε float64) bool {
	return math.Abs(a-b) <= ε
}
