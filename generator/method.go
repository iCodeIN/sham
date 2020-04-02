package generator

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

// generateMethod generates the code for a single method.
func generateMethod(
	out *jen.File,
	t *ast.TypeSpec,
	m *ast.Field,
) {
	out.Func().
		Params(
			jen.Id("x").
				Id("*" + t.Name.Name),
		).
		Id(m.Names[0].Name).
		ParamsFunc(func(grp *jen.Group) {
		}).
		BlockFunc(func(grp *jen.Group) {
			grp.If(jen.Id("x").Dot(fieldName(m)).Op("!=").Nil()).
				Block(
					jen.Id("x").Dot(fieldName(m)).Call(),
				)
			grp.Line()
			grp.If(jen.Id("x").Dot(t.Name.Name).Op("!=").Nil()).
				Block(
					jen.Id("x").Dot(t.Name.Name).Dot(m.Names[0].Name).Call(),
				)
		})
}
