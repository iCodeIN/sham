package inputs

type InputParam interface {
	Anon(int)
	Single(v int)
	Multiple(a int, b float64)
	MultipleNames(a, b int, c, d float64)
	Variadic(args ...int)
	ReceiverCollision(x int)
}
