package inputs

type InputParam interface {
	Scalar(v int)
	Anon(int)
	Multiple(a int, b float64)
	MultipleNames(a, b int)
}
