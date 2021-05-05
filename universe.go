// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

		p.Pos = p.Pos.Add(dir.Mul(ps.Velocity))
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

func (u *Universe) neighbours(pos, dir Vec2, i int) (L, R int) {
	r := u.ParamSet.Radius
	for j, p := range u.Particles {
		if i == j {
			continue
		}
		if withinRadius(pos, p.Pos, r) {
			if isLeft(pos, pos.Add(dir), p.Pos) {
				L++
			} else {
				R++
			}
		}
	}
	return
}

func isLeft(a, b, p Vec2) bool {
	return (p.X-a.X)*(b.Y-a.Y)-(p.Y-a.Y)*(b.X-a.X) > 0
}

func withinRadius(a, b Vec2, r float64) bool {
	return a.Dist(b) <= r
}

var (
	colorGreen   = color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff}
	colorBlue    = color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}
	colorBrown   = color.RGBA{R: 0x96, G: 0x4b, B: 0, A: 0xff}
	colorYellow  = color.RGBA{R: 0xff, G: 0xff, B: 0, A: 0xff}
	colorMagenta = color.RGBA{R: 0xff, G: 0, B: 0xff, A: 0xff}
)

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
