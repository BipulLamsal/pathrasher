package main

import "fmt"

type Color = Vector

func (c Color) RGB() (byte, byte, byte) {
	clamp := func(x float64) float64 {
		if x < 0 {
			return 0
		}
		if x > 0.999 {
			return 0.999
		}
		return x
	}

	r := byte(255.999 * clamp(c.X))
	g := byte(255.999 * clamp(c.Y))
	b := byte(255.999 * clamp(c.Z))

	return r, g, b
}

func (c Color) Add(o Color) Color {
	return Color{c.X + o.X, c.Y + o.Y, c.Z + o.Z}
}

func (c Color) MulScalar(t float64) Color {
	return Color{c.X * t, c.Y * t, c.Z * t}
}

func (c Color) MulColor(o Color) Color {
	return Color{c.X * o.X, c.Y * o.Y, c.Z * o.Z}
}
