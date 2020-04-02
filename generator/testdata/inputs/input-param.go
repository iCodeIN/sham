package inputs

type InputParam interface {
	Anon(int, string)
	Single(v int)
	Multiple(a int, b string)
	MultipleNames(a, b int, c, d string)
	Variadic(args ...int)
}
