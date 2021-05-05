// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Dist(w Vec2) float64 {
	return v.Sub(w).Len()
}

func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

func (v Vec2) Len() float64 {
	return math.Sqrt(v.SqLen())
}

func (v Vec2) SqLen() float64 {
	return v.Dot(v)
}

func (v Vec2) Dot(w Vec2) float64 {
	return v.X*w.X + v.Y*w.Y
}

func (v Vec2) Mul(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}
