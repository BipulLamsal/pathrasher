package geometry

import "pathrasher/ptmath"

type HitRecord struct {
	Point     ptmath.Vector // Intersection point
	Normal    ptmath.Vector // Surface normal at intersection
	T         float64       // Distance along ray
	FrontFace bool          // Did ray hit from outside?
}

// SetFaceNormal determines which side of the surface the ray hit
func (h *HitRecord) SetFaceNormal(ray *Ray, outwardNormal ptmath.Vector) {
	h.FrontFace = ray.Direction.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Mul(-1) // Flip normal if hit from inside
	}
}

/*
	If a ray bounces off of a material and keeps 100% of its color,

then we say that the material is white. If a ray bounces off of a
material and keeps 0% of its color, then we say that the material is black.
As a first demonstration of our new diffuse material we'll set the
ray_color function to return 50 % of the color from a bounce.
We should expect to get a nice gray color.
*/
func (h *HitRecord) RandomOn() ptmath.Vector {
	randomUnitSphere := ptmath.RandomNormal()
	if randomUnitSphere.Dot(h.Normal) > 0.0 { // In the same hemisphere as the normal
		return randomUnitSphere
	} else {
		return randomUnitSphere.Mul(-1)
	}
}

// Hittable interface - any shape that can be intersected by a ray
type Hittable interface {
	Hit(ray *Ray, tMin, tMax float64, rec *HitRecord) bool
}
