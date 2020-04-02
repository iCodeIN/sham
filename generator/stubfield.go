package generator

import (
	"go/ast"

	"github.com/dave/jennifer/jen"
)

func (v *visitor) generateStubField(out *jen.Group) {
	out.Commentf(
		"%s is an implementation of the %s() method.",
		v.fieldName,
		v.methodName,
	)
	out.Commentf(
		"If it is non-nil, it takes precedence over the embedded %s interface.",
		v.interfaceName,
	)
	out.Id(v.fieldName).
		Func().
		ParamsFunc(
			func(out *jen.Group) {
				v.generateTypeList(out, v.methodInputs)
			},
		).
		ParamsFunc(
			func(out *jen.Group) {
				v.generateTypeList(out, v.methodOutputs)
			},
		)
}

// generateSignature generates type lists from field lists.
func (v *visitor) generateTypeList(out *jen.Group, fields []*ast.Field) {
	for _, f := range fields {
		n := len(f.Names)
		if n == 0 {
			n = 1
		}

		for i := 0; i < n; i++ {
			out.Add(v.newType(f.Type))
		}
	}
}
