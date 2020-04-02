package inputs

type Collision interface {
	Input(stub int)
	Output() (stub int)
}
