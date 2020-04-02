package inputs

type OutputParam interface {
	Anon() int
	Single() (a int)
	Multiple() (a int, b float64)
	MultipleNames() (a, b int, c, d float64)
	ReceiverCollision() (x int)
}
