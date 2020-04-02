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
	n := fieldName(m)

	out.Commentf(
		"%s is an implementation of the %s() method.",
		n,
		m.Names[0].Name,
	)
	out.Commentf(
		"If it is non-nil, it takes precedence over x.%s.%s().",
		t.Name.Name,
		m.Names[0].Name,
	)
	out.Id(n).
		Func().
		ParamsFunc(
			func(grp *jen.Group) {
				generateSignature(grp, m)
			},
		)
}

// fieldName returns the name of the stub to use for the given method.
func fieldName(n *ast.Field) string {
	return n.Names[0].Name + "Func"
}
