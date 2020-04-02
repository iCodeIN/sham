package generator

import (
	"fmt"
	"go/ast"

	"github.com/dave/jennifer/jen"
)

func newParamType(
	n ast.Node,
) jen.Code {
	switch pt := n.(type) {
	case *ast.Ident:
		return jen.Id(pt.Name)
	case *ast.Ellipsis:
		return jen.Op("...").Add(
			newParamType(pt.Elt),
		)
	}

	panic(fmt.Sprintf("unsupported: %T", n))
}

// generateParams generates an parameter list.
func generateParams(
	out *jen.Group,
	m *ast.Field,
) {
	nt := nameTable{prefix: "i"}

	for _, p := range m.Type.(*ast.FuncType).Params.List {
		names := p.Names
		if len(names) == 0 {
			names = []*ast.Ident{{Name: ""}}
		}

		out.ListFunc(
			func(grp *jen.Group) {
				for _, name := range names {
					grp.Id(nt.Get(name.Name))
				}
			},
		).Add(
			newParamType(p.Type),
		)
	}
}

// generateSignature generates a parameter list with no variable names.
func generateSignature(
	out *jen.Group,
	m *ast.Field,
) {
	for _, p := range m.Type.(*ast.FuncType).Params.List {
		n := len(p.Names)
		if n == 0 {
			n = 1
		}

		for i := 0; i < n; i++ {
			out.Add(newParamType(p.Type))
		}
	}
}

// generateArgs generates an argument list.
func generateArgs(
	out *jen.Group,
	m *ast.Field,
) {
	nt := nameTable{prefix: "i"}

	for _, p := range m.Type.(*ast.FuncType).Params.List {
		names := p.Names
		if len(names) == 0 {
			names = []*ast.Ident{{Name: ""}}
		}

		if _, ok := p.Type.(*ast.Ellipsis); ok {
			out.Id(nt.Get(names[0].Name)).Op("...")
		} else {
			for _, name := range names {
				out.Id(nt.Get(name.Name))
			}
		}
	}
}
