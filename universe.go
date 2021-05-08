// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pps is an implementation of the Primordial Particle System (PPS)
// described in: Schmickl, T., Stefanec, M. & Crailsheim, K.
// How a life-like system emerges from a simplistic particle motion law.
// Sci Rep 6, 37969 (2016). https://doi.org/10.1038/srep37969
package pps

import "image/color"

type Universe struct {
	Size      Vec2
	ParamSet  ParamSet
	Particles []Particle
}

type ParamSet struct {
	Alpha    float64
	Beta     float64
	Velocity float64
	Radius   float64
}

func NewUniverse(size Vec2, particles int, ps ParamSet) *Universe {
	u := &Universe{
		Size:      size,
		ParamSet:  ps,
		Particles: make([]Particle, particles),
	}
	for i := range u.Particles {
		u.Particles[i] = randomParticle(size)
	}
	return u
}

func (u *Universe) Step() {
	ps := u.ParamSet
	for i, p := range u.Particles {
		dir := p.Dir()

		p.Pos = p.Pos.add(dir.mul(ps.Velocity))
		p.Pos.X = wrap(p.Pos.X, u.Size.X)
		p.Pos.Y = wrap(p.Pos.Y, u.Size.Y)

		L, R := u.neighbours(p.Pos, dir, i)
		N := L + R
		deltaPhi := ps.Alpha + ps.Beta*float64(N*sign(R-L))
		p.Angle += deltaPhi
		for p.Angle <= -180 {
			p.Angle += 360
		}
		for p.Angle > 180 {
			p.Angle -= 360
		}

		p.Color = particleColor(N)

		u.Particles[i] = p
	}
}

// neighbours returns the number of neighbours on the left side (L) and on the
// right side (R) of a particle at position 'pos' with direction 'dir'.
// The left side and the right side are semicircles with the radius from the
// universe's parameter set. The sum N=L+R is the total number of neighbours
// within this radius. Parameter i is the index of the particle in the
// universe's Particles slice.
func (u *Universe) neighbours(pos, dir Vec2, i int) (L, R int) {
	r := u.ParamSet.Radius
	for j, p := range u.Particles {
		if i == j {
			continue
		}
		if withinRadius(pos, p.Pos, r) {
			if isLeft(p.Pos, pos, pos.add(dir)) {
				L++
			} else {
				R++
			}
		}
	}
	return
}

// isLeft reports whether point p is to the left side of the line
// through a and b.
func isLeft(p, a, b Vec2) bool {
	return (p.X-a.X)*(b.Y-a.Y)-(p.Y-a.Y)*(b.X-a.X) > 0
}

func withinRadius(a, b Vec2, r float64) bool {
	return a.dist(b) <= r
}

var (
	colorGreen   = color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff}
	colorBlue    = color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}
	colorBrown   = color.RGBA{R: 0x96, G: 0x4b, B: 0, A: 0xff}
	colorYellow  = color.RGBA{R: 0xff, G: 0xff, B: 0, A: 0xff}
	colorMagenta = color.RGBA{R: 0xff, G: 0, B: 0xff, A: 0xff}
)

// particleColor returns the color for a particle with N neighbours.
func particleColor(N int) color.Color {
	if N < 13 {
		return colorGreen
	}
	if N <= 15 {
		return colorBrown
	}
	if N <= 35 {
		return colorBlue
	}
	return colorYellow
}
