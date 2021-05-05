// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) dist(w Vec2) float64 {
	return v.sub(w).len()
}

func (v Vec2) add(w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

func (v Vec2) sub(w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

func (v Vec2) len() float64 {
	return math.Sqrt(v.sqLen())
}

func (v Vec2) sqLen() float64 {
	return v.dot(v)
}

func (v Vec2) dot(w Vec2) float64 {
	return v.X*w.X + v.Y*w.Y
}

func (v Vec2) mul(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}
