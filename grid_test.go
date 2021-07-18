// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "testing"

func TestIsLeft(t *testing.T) {
	tests := []struct {
		a, b, p Vec2
		want    bool
	}{
		{p: Vec2{X: 5, Y: -1}, a: Vec2{X: 0, Y: 0}, b: Vec2{X: 10, Y: 0}, want: true},
		{p: Vec2{X: 5, Y: 0}, a: Vec2{X: 0, Y: 0}, b: Vec2{X: 10, Y: 0}, want: false},
		{p: Vec2{X: 5, Y: 1}, a: Vec2{X: 0, Y: 0}, b: Vec2{X: 10, Y: 0}, want: false},
		{p: Vec2{X: 11, Y: 15}, a: Vec2{X: 10, Y: 20}, b: Vec2{X: 15, Y: 10}, want: true},
		{p: Vec2{X: 17, Y: 12}, a: Vec2{X: 10, Y: 20}, b: Vec2{X: 15, Y: 10}, want: false},
		{p: Vec2{X: 5, Y: 1}, a: Vec2{X: 0, Y: 0}, b: Vec2{X: -10, Y: 0}, want: true},
		{p: Vec2{X: 5, Y: -1}, a: Vec2{X: 0, Y: 0}, b: Vec2{X: -10, Y: 0}, want: false},
	}
	for _, tt := range tests {
		got := isLeft(tt.p, tt.a, tt.b)
		if got != tt.want {
			t.Errorf("isLeft(p: %v, a: %v, b: %v) = %v, want %v",
				tt.p, tt.a, tt.b, got, tt.want)
		}
	}
}

func TestWithinRadius(t *testing.T) {
	tests := []struct {
		a, b Vec2
		r    float64
		want bool
	}{
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 80}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 82}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 85}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 85.1}, r: 5, want: false},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 75}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{100, 74.9}, r: 5, want: false},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{99, 80}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{95, 80}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{94.9, 80}, r: 5, want: false},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{104, 80}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{105, 80}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{105.1, 80}, r: 5, want: false},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{103, 83}, r: 5, want: true},
		{a: Vec2{X: 100, Y: 80}, b: Vec2{103.6, 83.6}, r: 5, want: false},

		{a: Vec2{X: 50, Y: 70}, b: Vec2{50, 60}, r: 10, want: true},
		{a: Vec2{X: 50, Y: 60}, b: Vec2{50, 70}, r: 10, want: true},
		{a: Vec2{X: 40, Y: 60}, b: Vec2{50, 60}, r: 10, want: true},
		{a: Vec2{X: 60, Y: 60}, b: Vec2{50, 60}, r: 10, want: true},
		{a: Vec2{X: 50, Y: 70}, b: Vec2{50, 59.9}, r: 10, want: false},
	}
	for _, tt := range tests {
		got := withinRadius(tt.a, tt.b, tt.r)
		if got != tt.want {
			t.Errorf("withinRadius(a: %v, b: %v, r: %g) = %v, want %v",
				tt.a, tt.b, tt.r, got, tt.want)
		}
	}
}
