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

// generateInputParams generates an input parameter list.
func generateInputParams(
	out *jen.Group,
	list *ast.FieldList,
) {
	if list == nil {
		return
	}

	nt := nameTable{prefix: "i"}

	for _, p := range list.List {
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

// generateOutputParams generates an output parameter list.
func generateOutputParams(
	out *jen.Group,
	list *ast.FieldList,
) {
	if list == nil || len(list.List) == 0 {
		return
	}

	anon := true
	for _, p := range list.List {
		if len(p.Names) > 0 {
			anon = false
			break
		}
	}

	if anon {
		generateSignature(
			out,
			list,
		)

		return
	}

	// dapper.Print(list)

	panic("ni")
}

// generateSignature generates a parameter list with no variable names.
func generateSignature(
	out *jen.Group,
	list *ast.FieldList,
) {
	if list == nil {
		return
	}

	for _, p := range list.List {
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
	list *ast.FieldList,
) {
	if list == nil {
		return
	}

	nt := nameTable{prefix: "i"}

	for _, p := range list.List {
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
