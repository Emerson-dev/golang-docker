package calc

type IMath interface {
	Calc() int
}

type Sum struct {
	X int
	Y int
}

func (s Sum) Calc() int {
	return s.X + s.Y
}