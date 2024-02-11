// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "testing"

func BenchmarkUniverseStep100x100(b *testing.B) {
	benchmarkUniverseStep(b, Vec2{X: 100, Y: 100})
}

func BenchmarkUniverseStep250x250(b *testing.B) {
	benchmarkUniverseStep(b, Vec2{X: 250, Y: 250})
}

func BenchmarkUniverseStep400x400(b *testing.B) {
	benchmarkUniverseStep(b, Vec2{X: 400, Y: 400})
}

func BenchmarkUniverseStep800x800(b *testing.B) {
	benchmarkUniverseStep(b, Vec2{X: 800, Y: 800})
}

func benchmarkUniverseStep(b *testing.B, size Vec2) {
	params := ParamSet{
		Alpha:    180,
		Beta:     17,
		Velocity: 0.67,
		Radius:   5.0,
	}
	u := NewUniverse(size, int(size.X*size.Y*0.08), params)
	for range b.N {
		u.Step()
	}
}
