package generator

import (
	"fmt"
	"go/ast"

	"github.com/dave/jennifer/jen"
)

// generateMethod generates the code for a single method.
func generateMethod(
	out *jen.File,
	t *ast.TypeSpec,
	m *ast.Field,
) {
	ftype := m.Type.(*ast.FuncType)
	structName := t.Name.Name
	methodName := m.Names[0].Name
	stubName := fieldName(m)

	hasReturns := ftype.Results != nil &&
		len(ftype.Results.List) > 0

	generateCall := func(grp *jen.Group, fn jen.Code) {
		var stmt *jen.Statement

		if hasReturns {
			stmt = grp.Return().Add(fn)
		} else {
			stmt = grp.Add(fn)
		}

		stmt.CallFunc(
			func(grp *jen.Group) {
				generateArgs(
					grp,
					ftype.Params,
				)
			},
		)
	}

	out.Func().
		Params(
			jen.Id("x").
				Id("*" + structName),
		).
		Id(methodName).
		ParamsFunc(
			func(grp *jen.Group) {
				generateInputParams(
					grp,
					ftype.Params,
				)
			},
		).
		ListFunc(
			func(grp *jen.Group) {
				generateSignature(grp, ftype.Results)
			},
		).
		BlockFunc(
			func(grp *jen.Group) {
				grp.If(jen.Id("x").Dot(stubName).Op("!=").Nil()).
					BlockFunc(
						func(grp *jen.Group) {
							generateCall(
								grp,
								jen.Id("x").Dot(stubName),
							)
						},
					)

				grp.Line()

				grp.If(jen.Id("x").Dot(structName).Op("!=").Nil()).
					BlockFunc(
						func(grp *jen.Group) {
							generateCall(
								grp,
								jen.Id("x").Dot(structName).Dot(methodName),
							)
						},
					)

				if hasReturns {
					grp.Line()

					grp.Panic(
						jen.Lit(
							fmt.Sprintf(
								"%s() has no implementation, set the %s or %s field",
								methodName,
								structName,
								stubName,
							),
						),
					)
				}
			},
		)
}
