package circle

import "math"

type Circle struct {
	Radio float64
}

func (c Circle) Area() float64 {
	pi := math.Pi
	return pi * math.Pow(c.Radio, 2)
}
