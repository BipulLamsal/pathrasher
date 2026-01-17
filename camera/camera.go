package camera

import (
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"os"
	"pathrasher/color"
	"pathrasher/geometry"
	"pathrasher/ptmath"
)

type Camera struct {
	AspectRatio     float64
	ImageWidth      int
	SamplesPerPixel int

	imageHeight int
	center      ptmath.Vector
	pixel00Loc  ptmath.Vector
	pixelDeltaU ptmath.Vector
	pixelDeltaV ptmath.Vector
}

func (c *Camera) Initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	c.imageHeight = max(c.imageHeight, 1)

	fmt.Println("Image Heigth", c.imageHeight)
	fmt.Println("Image Width", c.ImageWidth)

	if c.SamplesPerPixel <= 0 {
		c.SamplesPerPixel = 10
	}

	c.center = ptmath.Vector{X: 0, Y: 0, Z: 0}

	focalLength := 2.0
	viewportHeight := 5.0
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	fmt.Println("Viewport Height", viewportHeight)
	fmt.Println("Viewport Width", viewportWidth)

	viewportU := ptmath.Vector{X: viewportWidth, Y: 0, Z: 0}
	viewportV := ptmath.Vector{X: 0, Y: -viewportHeight, Z: 0}

	c.pixelDeltaU = viewportU.Mul(1.0 / float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.Mul(1.0 / float64(c.imageHeight))

	viewportUpperLeft := c.center.
		Sub(ptmath.Vector{X: 0, Y: 0, Z: focalLength}).
		Sub(viewportU.Mul(0.5)).
		Sub(viewportV.Mul(0.5))

	fmt.Println(viewportUpperLeft)

	c.pixel00Loc = viewportUpperLeft.Add(c.pixelDeltaU.Add(c.pixelDeltaV).Mul(0.5))
	fmt.Println(c.pixel00Loc)
}

func (c *Camera) Render(out io.Writer, world geometry.Hittable) {
	fmt.Fprintf(out, "P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight)

	for j := 0; j < c.imageHeight; j++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", c.imageHeight-j)
		for i := 0; i < c.ImageWidth; i++ {
			pixelColor := color.Color{}
			pixelCenter := c.pixel00Loc.
				Add(c.pixelDeltaU.Mul(float64(i))).
				Add(c.pixelDeltaV.Mul(float64(j)))

			for s := 0; s < c.SamplesPerPixel; s++ {
				uOffset := rand.Float64()
				vOffset := rand.Float64()
				samplePoint := pixelCenter.
					Add(c.pixelDeltaU.Mul(uOffset)).
					Add(c.pixelDeltaV.Mul(vOffset))
				rayDirection := samplePoint.Sub(c.center)
				r := geometry.Ray{Origin: c.center, Direction: rayDirection}
				pixelColor.Add(rayColor(&r, world, 50))
			}
			pixelColor.MulScalar(1.0 / float64(c.SamplesPerPixel))

			color.WriteColor(out, pixelColor)
		}
	}
	fmt.Fprintln(os.Stderr, "\rDone.                 ")
}

func rayColor(r *geometry.Ray, world geometry.Hittable, depth int) color.Color {
	if depth <= 0 {
		return color.Color{0, 0, 0}
	}
	rec := geometry.HitRecord{}
	if world.Hit(r, 0, math.Inf(1), &rec) {
		direction := rec.RandomOn()
		result := rayColor(&geometry.Ray{Origin: rec.Point, Direction: direction}, world, depth-1)
		result.MulScalar(0.5)
		return result
		// n := rec.Normal
		// // 0.5 * (rec.Normal + 1.0)
		// return color.Color{
		// 	R: 0.35 * (n.X + 1),
		// 	G: 0.55 * (n.Y + 1),
		// 	B: 0.30 * (n.Z + 1),
		// }
	}

	unitDirection := r.Direction.Normalize()
	a := 0.5 * (unitDirection.Y + 1.0) // Fade from the bottom to top

	startColor := color.Color{R: 1.0, G: 1.0, B: 1.0}
	startColor.MulScalar(1.0 - a)

	endColor := color.Color{R: 0.5, G: 0.7, B: 1.0}
	endColor.MulScalar(a)

	startColor.Add(endColor)
	return startColor
}
