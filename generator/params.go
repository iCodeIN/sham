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
	for _, p := range m.Type.(*ast.FuncType).Params.List {
		id := out.Id(p.Names[0].Name)

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
	for _, p := range m.Type.(*ast.FuncType).Params.List {
		out.Id(p.Names[0].Name)
	}
}
