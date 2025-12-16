package main

import (
	"fmt"
	"os"
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
		width, height = 1080,1080
		maxColorValue = 255
	)

	// Frame Buffer
	frameBuffer := make([]byte, width * height * 3) // 3 for R G B

	file, err := os.Create("output.ppm")
	check(err, "Open File")

	// since we have opend the file we need to defer the close 
	defer file.Close()

	// PPM requries its header to be written using in the file 
	// so we use Fprintf to write in the file because
	// file implements io.Writer 

	fmt.Fprintf(file, "P6\n%d %d\n%d\n", width, height, maxColorValue)
	for j := 0; j < height ; j++ {
		for i :=0 ; i < width; i++ {
			// u := float64(i) / float64(width-1)
			// v := float64(j) / float64(height-1)

			// r:= byte(maxColorValue * u)   
			// g:= byte(maxColorValue * v) 
			// b:= byte(maxColorValue * 0.2)

// type Sphere struct {
// 	Center Vector
// 	Radius float64
// }
			

			index := (j*width + i) * 3
			frameBuffer[index] = r
			frameBuffer[index+1] = g
			frameBuffer[index+2] = b
			np := int((float64(index+3) / float64(cap(frameBuffer))) * 100)

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
