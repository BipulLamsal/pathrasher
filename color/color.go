package color 

type Color struct {
	R, G, B float64
}

var (
 Gray = Color{0.5, 0.5, 0.5}
 Red = Color{1.0, 0.0, 0.0}
 Blue = Color{0.0, 0.0, 1.0}
 )

func (c *Color) RGB() (byte, byte, byte) {
	clamp := func(x float64) float64 {
		if x < 0 {
			return 0
		}
		if x > 0.999 {
			return 0.999
		}
		return x
	}

	r := byte(255.999 * clamp(c.R))
	g := byte(255.999 * clamp(c.G))
	b := byte(255.999 * clamp(c.B))

	return r, g, b
}


func (c *Color) Add(o Color) {
	c.R += o.R
	c.G += o.G
	c.B += o.B
}

func (c *Color) Normalize(){
	c.R /= 255.0
	c.G /= 255.0
	c.B /= 255.0
}

func (c *Color) MulScalar(t float64) {
	c.R *= t
	c.G *= t
	c.B *= t
}

func (c *Color) MulColor(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}
