// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pps is an implementation of the Primordial Particle System (PPS)
// described in: Schmickl, T., Stefanec, M. & Crailsheim, K.
// How a life-like system emerges from a simplistic particle motion law.
// Sci Rep 6, 37969 (2016). https://doi.org/10.1038/srep37969
package pps

import (
	"image/color"
	"math"
)

type Universe struct {
	Size      Vec2
	ParamSet  ParamSet
	Particles []Particle
	grid      particleGrid
}

type ParamSet struct {
	Alpha    float64
	Beta     float64
	Velocity float64
	Radius   float64
}

func NewUniverse(size Vec2, particleCount int, ps ParamSet) *Universe {
	particles := make([]Particle, particleCount)
	cellSize := int(math.Ceil(ps.Radius))
	grid := makeParticleGrid(cellSize)
	for i := 0; i < particleCount; i++ {
		particles[i] = randomParticle(size)
		grid.addParticle(&particles[i])
	}
	u := &Universe{
		Size:      size,
		ParamSet:  ps,
		Particles: particles,
		grid:      grid,
	}
	return u
}

func (u *Universe) Step() {
	ps := u.ParamSet
	for i := range u.Particles {
		p := &u.Particles[i]
		dir := p.Dir()

		pos := p.Pos.add(dir.mul(ps.Velocity))
		pos.X = wrap(pos.X, u.Size.X)
		pos.Y = wrap(pos.Y, u.Size.Y)
		u.grid.updatePos(p, pos)

		L, R, NClose := u.grid.neighbours(p, ps.Radius, 1.3)
		N := L + R
		deltaPhi := ps.Alpha + ps.Beta*float64(N*sign(R-L))
		p.Angle += deltaPhi
		for p.Angle <= -180 {
			p.Angle += 360
		}
		for p.Angle > 180 {
			p.Angle -= 360
		}

		p.Color = particleColor(N, NClose)
	}
}

var (
	colorGreen   = color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff}
	colorBlue    = color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}
	colorBrown   = color.RGBA{R: 0x96, G: 0x4b, B: 0, A: 0xff}
	colorYellow  = color.RGBA{R: 0xff, G: 0xff, B: 0, A: 0xff}
	colorMagenta = color.RGBA{R: 0xff, G: 0, B: 0xff, A: 0xff}
)

// particleColor returns the color for a particle with N neighbours,
// of which NClose are close neighbours.
func particleColor(N, NClose int) color.Color {
	if NClose > 15 {
		return colorMagenta
	}
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
