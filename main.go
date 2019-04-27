package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Scene struct {
	Width  uint
	Height uint
	Fov    float64
	Sphere Sphere
}

func main() {
	scene := Scene{
		Width:  800,
		Height: 600,
		Fov:    90.0,
		Sphere: Sphere{
			Center: Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -5.0,
			},
			Radius: 1.0,
			Color:  color.RGBA{50, 40, 150, 255},
		},
	}

	img := render(&scene)

	outputFile, err := os.Create("step1.png")
	if err != nil {
		panic("Could not create file")
	}
	defer outputFile.Close()

	png.Encode(outputFile, img)
}

func render(scene *Scene) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(scene.Width), int(scene.Height)))
	count := 0
	for x := uint(0); x < scene.Width; x++ {
		for y := uint(0); y < scene.Height; y++ {
			ray := CreatePrime(x, y, scene)
			count++

			if scene.Sphere.intersect(&ray) {
				img.Set(int(x), int(y), scene.Sphere.Color)
			} else {
				img.Set(int(x), int(y), color.Black)
			}
		}
	}

	log.Printf("Done!")
	return img
}
