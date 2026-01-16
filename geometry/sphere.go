/*
https://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/ray-sphere-intersection.html
*/
package geometry

import (
	"math"
	ptmath "pathrasher/ptmath"
)

type Sphere struct {
	Center ptmath.Vector
	Radius float64
}

// to check the intersection of the line and the Sphere
// If the discriminant is:
//
//	positive - there are 2 real solutions
//	negative - there are 0 real solutions
//	zero - there is 1 real solution
func (s *Sphere) Hit(ray *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.Dot(ray.Direction)
	halfB := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtd) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrtd) / a
		if root < tMin || root > tMax {
			return false
		}
	}

	rec.T = root
	rec.Point = ray.At(root)
	// N = (intersection - centerofsphere) * 1/radius
	outwardNormal := rec.Point.Sub(s.Center).Mul(1 / float64(s.Radius))
	rec.SetFaceNormal(ray, outwardNormal)

	return true
}
