package geometry

import "pathrasher/ptmath"

type HitRecord struct {
    Point  ptmath.Vector // Intersection point
    Normal ptmath.Vector // Surface normal at intersection
    T      float64       // Distance along ray
    FrontFace bool       // Did ray hit from outside?
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

// Hittable interface - any shape that can be intersected by a ray
type Hittable interface {
    Hit(ray *Ray, tMin, tMax float64, rec *HitRecord) bool
}
