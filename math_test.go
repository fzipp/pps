// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "testing"

func TestSign(t *testing.T) {
	tests := []struct {
		i    int
		want int
	}{
		{i: -1234, want: -1},
		{i: -100, want: -1},
		{i: -5, want: -1},
		{i: -2, want: -1},
		{i: -1, want: -1},
		{i: 0, want: 0},
		{i: 1, want: 1},
		{i: 2, want: 1},
		{i: 5, want: 1},
		{i: 100, want: 1},
		{i: 1234, want: 1},
	}
	for _, tt := range tests {
		got := sign(tt.i)
		if got != tt.want {
			t.Errorf("sign(%d) = %d, want %d", tt.i, got, tt.want)
		}
	}
}

func TestDeg2rad(t *testing.T) {
	tests := []struct {
		deg  float64
		want float64
	}{
		{deg: 0, want: 0},
		{deg: 45, want: 0.7853981633974483},
		{deg: 90, want: 1.5707963267948966},
		{deg: 135, want: 2.356194490192345},
		{deg: 180, want: 3.141592653589793},
		{deg: 225, want: 3.9269908169872414},
		{deg: 270, want: 4.71238898038469},
		{deg: 315, want: 5.497787143782138},
		{deg: 360, want: 6.283185307179586},
	}
	for _, tt := range tests {
		got := deg2rad(tt.deg)
		if got != tt.want {
			t.Errorf("deg2rad(%g) = %g, want %g", tt.deg, got, tt.want)
		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		n    float64
		max  float64
		want float64
	}{
		{n: 0, max: 100, want: 0},
		{n: -1, max: 100, want: 99},
		{n: 255, max: 250, want: 5},
		{n: 300, max: 250, want: 50},
		{n: -2.5, max: 250, want: 247.5},
	}
	for _, tt := range tests {
		got := wrap(tt.n, tt.max)
		if got != tt.want {
			t.Errorf("wrap(%g, %g) = %g, want %g", tt.n, tt.max, got, tt.want)
		}
	}
}
