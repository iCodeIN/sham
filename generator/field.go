package generator

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

// generateField generates a "XXXFunc" field for single method.
func generateField(
	out *jen.Group,
	t *ast.TypeSpec,
	m *ast.Field,
) {
	ftype := m.Type.(*ast.FuncType)
	methodName := m.Names[0].Name
	stubName := fieldName(m)

	out.Commentf(
		"%s is an implementation of the %s() method.",
		stubName,
		methodName,
	)
	out.Commentf(
		"If it is non-nil, it takes precedence over x.%s.%s().",
		t.Name.Name,
		methodName,
	)
	out.Id(stubName).
		Func().
		ParamsFunc(
			func(grp *jen.Group) {
				generateSignature(grp, ftype.Params)
			},
		).
		ListFunc(
			func(grp *jen.Group) {
				generateSignature(grp, ftype.Results)
			},
		)
}

// fieldName returns the name of the stub to use for the given method.
func fieldName(n *ast.Field) string {
	return n.Names[0].Name + "Func"
}
