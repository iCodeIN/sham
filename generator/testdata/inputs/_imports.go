package inputs

import "io"

type Imports interface {
	Method(w io.Writer)
}
