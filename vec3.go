package main

import "math"

// Vec3 represents a 3D vector
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Length returns the length, or magnitude of a vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize converts a vector in to a unit vector
func (v *Vec3) Normalize() {
	length := v.Length()

	v.X = v.X / length
	v.Y = v.Y / length
	v.Z = v.Z / length
}

// Dot returns the dot product of 2 vectors
func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}
