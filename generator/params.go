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
		var name string

		if len(p.Names) == 0 {
			name = nt.Get("")
		} else {
			name = nt.Get(p.Names[0].Name)
		}

		id := out.Id(name)

		switch pt := p.Type.(type) {
		case *ast.Ident:
			id.Id(pt.Name)
		}
	}
}

// generateSignature generates a parameter list with no variable names.
func generateSignature(
	out *jen.Group,
	m *ast.Field,
) {
	for _, p := range m.Type.(*ast.FuncType).Params.List {
		switch pt := p.Type.(type) {
		case *ast.Ident:
			out.Id(pt.Name)
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
		var name string

		if len(p.Names) == 0 {
			name = nt.Get("")
		} else {
			name = nt.Get(p.Names[0].Name)
		}

		out.Id(name)
	}
}
