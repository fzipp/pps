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
