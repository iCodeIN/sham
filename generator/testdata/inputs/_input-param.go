package inputs

import "io"

type InputParam interface {
	Method(w io.Writer)
}
