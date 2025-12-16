/*

https://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/ray-sphere-intersection.html 

*/
package main
import (
	"math"
)

type Sphere struct {
	Center Vector
	Radius float64
}
// to check the intersection of the line and the Sphere
// If the discriminant is:
//
//     positive - there are 2 real solutions
//     negative - there are 0 real solutions
//     zero - there is 1 real solution

func (s Sphere) Hit(r Ray) (bool, float64) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false, 0
	}

	t := (-b - math.Sqrt(discriminant)) / (2 * a)
	if t < 0 {
		t = (-b + math.Sqrt(discriminant)) / (2 * a)
		if t < 0 {
			return false, 0
		}
	}

	// t is the distance that tells how far along the ray from camera      
	return true, t
}





