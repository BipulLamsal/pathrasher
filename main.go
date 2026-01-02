package main
import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"pathrasher/color"
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
	// Image
	const (
		width, height = 1080, 1080
		maxColorValue = 255
	)
	const aliasingSamples = 100
	
	// Frame Buffer
	frameBuffer := make([]byte, width * height * 3) // 3 for R G B
	file, err := os.Create("output.ppm")
	check(err, "Open File")
	
	// since we have opend the file we need to defer the close 
	defer file.Close()
	
	// Create world (hittable list) and add sphere to it
	world := geometry.World{}
	world.Add(&geometry.Sphere{
		Center: ptmath.Vector{X: 0, Y: 0, Z: 50},
		Radius: 10,
	})

	world.Add(&geometry.Sphere{
		Center: ptmath.Vector{X: 10, Y: 10, Z: 50},
		Radius: 10,
	})
	
	// PPM requries its header to be written using in the file 
	// so we use Fprintf to write in the file because
	// file implements io.Writer 
	fmt.Fprintf(file, "P6\n%d %d\n%d\n", width, height, maxColorValue)
	
	for j := range height {
		for i := range width {
			accumulatedColor := color.Color{}
			
			// Anti-aliasing: shoot multiple rays per pixel
			for range aliasingSamples {
				jitterU := (float64(i) + rand.Float64()) / float64(width)
				jitterV := (float64(j) + rand.Float64()) / float64(height)
				// map pixel to viewport [-1, 1]
				// 0 0 is exact center because of it  
				u := jitterU*2 - 1    // -1 left, +1 right
				v := 1 - jitterV*2    // -1 top, +1 bottom
				
				ray := geometry.Ray{
					Origin:    ptmath.Vector{X: 0, Y: 0, Z: -10},
					Direction: ptmath.Vector{X: u, Y: v, Z: 1}.Normalize(),
				}
				
				rec := geometry.HitRecord{}
				if world.Hit(&ray, 0.001, math.MaxFloat64, &rec) {
					// Hit the sphere - use a gray color
					accumulatedColor.Add(color.Color{R: 30.0/255.0, G: 30.0/255.0, B: 30.0/255.0})
				} else {
					// Sky color
					accumulatedColor.Add(color.Color{R: 135.0/255.0, G: 206.0/255.0, B: 235.0/255.0})
				}
			}
			
			// Average the samples
			accumulatedColor.MulScalar(1.0 / float64(aliasingSamples))
			
			index := (j*width + i) * 3
			r, g, b := accumulatedColor.RGB()
			frameBuffer[index] = r
			frameBuffer[index+1] = g
			frameBuffer[index+2] = b
			
			np := int((float64(index+3) / float64(len(frameBuffer))) * 100)
			dashes := [10]rune{}
			for i, j := np/10, 0; j < 10; j++ {
				if j <= i {
					dashes[j] = '/'
				} else {
					dashes[j] = ' '
				}
			}
			fmt.Printf("\r[%s] %d%%", string(dashes[:]), np)
		}
	}
	
	_, err = file.Write(frameBuffer)
	fmt.Printf("\n")
	check(err, "Write")
}
