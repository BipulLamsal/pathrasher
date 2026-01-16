// Ray is the concept that represents a line of sight or light path in 3D space
// It is a laser pointer shooting from the camera into the scene

// Point(t)=Origin+t⋅Direction t is greater or equals to zero // t=1 is one unit along the direction

package geometry 

import "pathrasher/ptmath"


type Ray struct {
	Origin, Direction ptmath.Vector
}

func (r Ray) At(t float64) ptmath.Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}

// we can impelenment refract, relect , transform matrix as well lets leave it for now


