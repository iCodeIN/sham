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
	var (
		funcType      = m.Type.(*ast.FuncType)
		structName    = t.Name.Name
		methodName    = m.Names[0].Name
		stubName      = fieldName(m)
		variableNames = newNameTable(funcType)
		anonNames     = map[int]string{}
		hasReturns    = funcType.Results != nil &&
			len(funcType.Results.List) > 0
	)

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
					anonNames,
					grp,
					funcType.Params,
				)
			},
		)
	}

	receiverName := variableNames.Get("x")

	out.Func().
		Params(
			jen.Id(receiverName).
				Id("*" + structName),
		).
		Id(methodName).
		ParamsFunc(
			func(grp *jen.Group) {
				generateInputParams(
					variableNames,
					anonNames,
					grp,
					funcType.Params,
				)
			},
		).
		// ListFunc(
		// 	func(grp *jen.Group) {
		// 		generateOutputParams(grp, funcType.Results)
		// 	},
		// ).
		BlockFunc(
			func(grp *jen.Group) {
				grp.If(jen.Id(receiverName).Dot(stubName).Op("!=").Nil()).
					BlockFunc(
						func(grp *jen.Group) {
							generateCall(
								grp,
								jen.Id(receiverName).Dot(stubName),
							)
						},
					)

				grp.Line()

				grp.If(jen.Id(receiverName).Dot(structName).Op("!=").Nil()).
					BlockFunc(
						func(grp *jen.Group) {
							generateCall(
								grp,
								jen.Id(receiverName).Dot(structName).Dot(methodName),
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
