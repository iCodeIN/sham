package inputs

import (
	"io"

	"golang.org/x/sync/errgroup"
	alias "golang.org/x/sync/singleflight"
)

type Imports interface {
	StdLib(r io.Reader) (w io.Writer)
	ThirdParty(g errgroup.Group)
	Aliased(g alias.Group)
}
