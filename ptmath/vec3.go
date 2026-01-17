package ptmath

import (
	"math"
	"math/rand"
)

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

func Random() Vector {
	return Vector{
		X: rand.Float64()*2.0 - 1.0,
		Y: rand.Float64()*2.0 - 1.0,
		Z: rand.Float64()*2.0 - 1.0,
	}
}

func RandomNormal() Vector {
	for true {
		vec := Random()
		length := vec.Length()
		if length <= 1 && length > math.SmallestNonzeroFloat64 {
			return vec.Normalize()
		}
	}
	return Vector{}
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
