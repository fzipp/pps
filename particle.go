// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import (
	"image/color"
	"math"
	"math/rand"
)

type Particle struct {
	Pos   Vec2
	Angle float64

	Color color.Color
}

// Dir returns the normed direction vector of the particle.
func (p *Particle) Dir() Vec2 {
	angleRad := deg2rad(p.Angle)
	return Vec2{
		X: math.Cos(angleRad),
		Y: math.Sin(angleRad),
	}
}

func randomParticle(screenSize Vec2) Particle {
	return Particle{
		Pos: Vec2{
			X: rand.Float64() * screenSize.X,
			Y: rand.Float64() * screenSize.Y,
		},
		Angle: (rand.Float64() * 360) - 180,
	}
}
