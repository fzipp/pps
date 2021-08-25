// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "testing"

func TestGridCellAdd(t *testing.T) {
	tests := []struct {
		c    gridCell
		x, y int
		want gridCell
	}{
		{c: gridCell{x: 3, y: 4}, x: 0, y: 0, want: gridCell{x: 3, y: 4}},
		{c: gridCell{x: 3, y: 4}, x: 1, y: 1, want: gridCell{x: 4, y: 5}},
		{c: gridCell{x: 3, y: 4}, x: -1, y: 0, want: gridCell{x: 2, y: 4}},
		{c: gridCell{x: 3, y: 4}, x: -1, y: -1, want: gridCell{x: 2, y: 3}},
		{c: gridCell{x: 6, y: 5}, x: 1, y: -1, want: gridCell{x: 7, y: 4}},
	}
	for _, tt := range tests {
		got := tt.c.add(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("%v.add(x: %d, y: %d) = %v, want %v",
				tt.c, tt.x, tt.y, got, tt.want)
		}
	}
}

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
