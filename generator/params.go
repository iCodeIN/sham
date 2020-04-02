package generator

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

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

		list := out.ListFunc(
			func(grp *jen.Group) {
				for _, name := range names {
					grp.Id(nt.Get(name.Name))
				}
			},
		)

		switch pt := p.Type.(type) {
		case *ast.Ident:
			list.Id(pt.Name)
		}
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
			switch pt := p.Type.(type) {
			case *ast.Ident:
				out.Id(pt.Name)
			}
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

		for _, name := range names {
			out.Id(nt.Get(name.Name))
		}
	}
}
