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
	var (
		funcType   = m.Type.(*ast.FuncType)
		typeName   = t.Name.Name
		methodName = m.Names[0].Name
		stubName   = fieldName(m)
	)

	out.Commentf(
		"%s is an implementation of the %s() method.",
		stubName,
		methodName,
	)
	out.Commentf(
		"If it is non-nil, it takes precedence over the embedded %s interface.",
		typeName,
	)
	out.Id(stubName).
		Func().
		ParamsFunc(
			func(grp *jen.Group) {
				generateSignature(grp, funcType.Params)
			},
		).
		ParamsFunc(
			func(grp *jen.Group) {
				generateSignature(grp, funcType.Results)
			},
		)
}

// fieldName returns the name of the stub to use for the given method.
func fieldName(n *ast.Field) string {
	return n.Names[0].Name + "Func"
}
