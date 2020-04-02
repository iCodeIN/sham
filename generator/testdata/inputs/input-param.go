package inputs

type InputParam interface {
	Scalar(v int)
	Anon(int)
}
