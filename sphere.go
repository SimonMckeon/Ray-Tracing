package main

import (
	"image/color"
)

type Sphere struct {
	Center Vec3
	Radius float64
	Color  color.Color
}

func (s *Sphere) intersect(r *Ray) bool {
	// Create a line segment between the ray origin and the center of the sphere
	segment := Vec3{
		s.Center.X - r.Origin.X,
		s.Center.Y - r.Origin.Y,
		s.Center.Z - r.Origin.Z,
	}

	// Use l as hypotenuse and find length of the adjacent side
	adj := segment.Dot(&r.Direction)

	// Find the length-squared of the opposite side
	d := segment.Dot(&segment) - (adj * adj)

	return d < s.Radius*s.Radius
}
