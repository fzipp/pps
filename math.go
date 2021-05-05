// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "math"

func sign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return +1
	}
	return 0
}

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