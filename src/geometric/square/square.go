package square

type Square struct {
	width  float64
	height float64
}

func (s Square) Area() float64 {
	return s.width * s.height
}
