// Ray is the concept that represents a line of sight or light path in 3D space
// It is a laser pointer shooting from the camera into the scene

// Point = Origin + t(scalar) * Direction // negative t means oppsoite to the origin
package geometry

import "pathrasher/ptmath"

type Ray struct {
	Origin, Direction ptmath.Vector
}

func (r Ray) At(t float64) ptmath.Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}

// we can impelenment refract, relect , transform matrix as well lets leave it for now
