package main

import (
	"fmt"
	"os"
	"pathrasher/camera"
	"pathrasher/geometry"
	"pathrasher/ptmath"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error(%s): %v\n", s, e)
		os.Exit(1)
	}
}

func main() {
	world := geometry.World{}
	world.Add(&geometry.Sphere{
		Center: ptmath.Vector{X: 0, Y: 0, Z: -3},
		Radius: 2,
	})
	world.Add(&geometry.Sphere{
		Center: ptmath.Vector{X: 0, Y: -120.5, Z: -20},
		Radius: 100,
	})

	cam := camera.Camera{
		AspectRatio:     16.0 / 9.0,
		ImageWidth:      400,
		SamplesPerPixel: 100,
	}
	cam.Initialize()

	file, err := os.Create("output.ppm")
	check(err, "Open File")
	defer file.Close()

	cam.Render(file, &world)
}
