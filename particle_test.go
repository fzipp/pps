// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "testing"

func TestParticleDir(t *testing.T) {
	tests := []struct {
		p    Particle
		want Vec2
	}{
		{p: Particle{Angle: 0}, want: Vec2{X: 1, Y: 0}},
		{p: Particle{Angle: 45}, want: Vec2{X: 0.70710678118, Y: 0.70710678118}},
		{p: Particle{Angle: 90}, want: Vec2{X: 0, Y: 1}},
		{p: Particle{Angle: 135}, want: Vec2{X: -0.70710678118, Y: 0.70710678118}},
		{p: Particle{Angle: 180}, want: Vec2{X: -1, Y: 0}},
		{p: Particle{Angle: -45}, want: Vec2{X: 0.70710678118, Y: -0.70710678118}},
		{p: Particle{Angle: -90}, want: Vec2{X: 0, Y: -1}},
		{p: Particle{Angle: -135}, want: Vec2{X: -0.70710678118, Y: -0.70710678118}},
		{p: Particle{Angle: -180}, want: Vec2{X: -1, Y: 0}},
	}
	for _, tt := range tests {
		got := tt.p.Dir()
		if !got.nearEq(tt.want) {
			t.Errorf("Particle with angle %g°: Dir() = %v, want: %v", tt.p.Angle, got, tt.want)
		}
	}
}
