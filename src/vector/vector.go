package vector

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Add(s Vector) {

	v.X += s.X
	v.Y += s.Y
}

func (v *Vector) Sub(s Vector) {

	v.X -= s.X
	v.Y -= s.Y
}

func (v *Vector) Mult(k float64) {

	v.X *= k
	v.Y *= k
}

func (v *Vector) Div(k float64) {

	if k > 0 {
		v.X /= k
		v.Y /= k
	}

}

func (v *Vector) Mod() float64 {

	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vector) Normalize() *Vector {

	if mod := v.Mod(); mod != 0 {

		v.X /= mod
		v.Y /= mod
	}

	return v
}
