package main

import "math"

// Intersectable defined an object that can intersect with a Ray
type Intersectable interface {
	intesect(r *Ray) bool
}

// A Ray represents a single Ray
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// CreatePrime creates a prime ray for a position in the viewport
func CreatePrime(x uint, y uint, scene *Scene) Ray {
	aspectRatio := float64(scene.Width) / float64(scene.Height)
	fovAdjustment := math.Tan(DegToRad(scene.Fov) / 2.0)

	// Add 0.5 because the ray should pass through the center of a pixel
	sensorXPixelCenter := float64(x) + 0.5
	sensorXNormalizedToWidth := sensorXPixelCenter / float64(scene.Width)
	sensorX := (sensorXNormalizedToWidth*2.0 - 1.0) * aspectRatio * fovAdjustment

	sensorYPixelCenter := float64(y) + 0.5
	sensorYNormalizedToWidth := sensorYPixelCenter / float64(scene.Height)
	sensorY := (1.0 - (sensorYNormalizedToWidth * 2.0)) * fovAdjustment

	r := Ray{
		Origin: Vec3{X: 0, Y: 0, Z: 0},
		// Z is -1 as rays move forward from camera
		Direction: Vec3{X: sensorX, Y: sensorY, Z: 1.0},
	}
	r.Normalize()

	return r
}

// Normalize normalizes the direction component of a Ray
func (r *Ray) Normalize() {
	r.Direction.Normalize()
}

// DegToRad converts degrees to radians
func DegToRad(d float64) float64 {
	return d * (math.Pi / 180)
}

// RadToDeg converts radians to degrees
func RadToDeg(r float64) float64 {
	return r * (180 / math.Pi)
}
