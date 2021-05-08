// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pps

import (
	"fmt"
	"math"
)

// A Vec2 represents a vector with coordinates X and Y in 2-dimensional
// euclidean space.
type Vec2 struct {
	X float64
	Y float64
}

// add returns the vector v+w.
func (v Vec2) add(w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

// sub returns the vector v-w.
func (v Vec2) sub(w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

// mul returns the vector v*s.
func (v Vec2) mul(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// dot returns the dot (a.k.a. scalar) product of v and w.
func (v Vec2) dot(w Vec2) float64 {
	return v.X*w.X + v.Y*w.Y
}

// dist returns the euclidean distance between two vectors.
func (v Vec2) dist(w Vec2) float64 {
	return v.sub(w).len()
}

// sqLen returns the square of the length (euclidean norm) of a vector.
func (v Vec2) sqLen() float64 {
	return v.dot(v)
}

// len returns the length (euclidean norm) of a vector.
func (v Vec2) len() float64 {
	return math.Sqrt(v.sqLen())
}

// NearEq returns whether v and w are approximately equal. This relation is not
// transitive in general. The tolerance for the floating-point components is
// ±1e-10.
func (v Vec2) nearEq(w Vec2) bool {
	return nearEq(v.X, w.X, epsilon) && nearEq(v.Y, w.Y, epsilon)
}

// String returns a string representation of v like "(3.25, -1.5)".
func (v Vec2) String() string {
	return fmt.Sprintf("(%g, %g)", v.X, v.Y)
}
