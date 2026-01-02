package ptmath 

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Add(u Vector) Vector {
	return Vector{
		v.X + u.X,
		v.Y + u.Y,
		v.Z + u.Z,
	}
}

func (v Vector) Sub(u Vector) Vector {
	return Vector{
		v.X - u.X,
		v.Y - u.Y,
		v.Z - u.Z,
	}
}

func (v Vector) Mul(s float64) Vector {
	return Vector{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}

func (v Vector) Dot(u Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vector) Cross(u Vector) Vector {
	return Vector{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Normalize() Vector {
	len := v.Length()
	if len == 0 {
		return Vector{0, 0, 0}
	}
	return v.Mul(1 / len)
}

