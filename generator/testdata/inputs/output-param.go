package inputs

type OutputParam interface {
	Anon() (int, string)
	Single() (a int)
	Multiple() (a int, b string)
	MultipleNames() (a, b int, c, d string)
}
