package override

const (
	PI = 3.14
)

type Form interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radio float64
}

func (c Circle) Area() float64 {
	return PI * c.Radio * c.Radio
}

type CompoundForm struct {
	Form
}
