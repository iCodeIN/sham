package generator

import (
	"fmt"
	"go/ast"

	"github.com/dave/jennifer/jen"
)

func (v *visitor) generateStubMethod(out *jen.Group) {
	out.Func().
		Params(
			jen.Id(v.vars.Receiver).
				Op("*").
				Id(v.structName),
		).
		Id(v.methodName).
		ParamsFunc(
			func(out *jen.Group) {
				v.generateInputParams(out)
			},
		).
		ParamsFunc(
			func(out *jen.Group) {
				v.generateOutputParams(out)
			},
		).
		BlockFunc(
			func(out *jen.Group) {
				out.
					If(
						jen.Id(v.vars.Receiver).Dot(v.fieldName).
							Op("!=").
							Nil(),
					).
					BlockFunc(
						func(out *jen.Group) {
							fn := jen.Id(v.vars.Receiver).Dot(v.fieldName)
							v.generateCall(out, fn)
						},
					)

				out.Line()

				out.
					If(
						jen.Id(v.vars.Receiver).Dot(v.interfaceName).
							Op("!=").
							Nil(),
					).
					BlockFunc(
						func(out *jen.Group) {
							fn := jen.Id(v.vars.Receiver).Dot(v.interfaceName).Dot(v.methodName)
							v.generateCall(out, fn)
						},
					)

				if len(v.methodOutputs) != 0 {
					out.Line()

					out.Panic(
						jen.Lit(
							fmt.Sprintf(
								"%s() has no implementation, set the %s or %s field",
								v.methodName,
								v.interfaceName,
								v.fieldName,
							),
						),
					)
				}
			},
		)
}

func (v *visitor) generateCall(
	out *jen.Group,
	fn *jen.Statement,
) {
	call := fn.CallFunc(
		func(out *jen.Group) {
			for i, f := range v.methodInputs {
				for _, name := range v.vars.Inputs[i] {
					id := out.Id(name)

					if _, ok := f.Type.(*ast.Ellipsis); ok {
						id.Op("...")
					}
				}
			}
		},
	)

	if len(v.methodOutputs) > 0 {
		out.Return(call)
	} else {
		out.Add(call)
	}
}

func (v *visitor) generateInputParams(out *jen.Group) {
	for i, f := range v.methodInputs {
		out.ListFunc(
			func(out *jen.Group) {
				for _, name := range v.vars.Inputs[i] {
					out.Id(name)
				}
			},
		).Add(v.newType(f.Type))
	}
}

func (v *visitor) generateOutputParams(out *jen.Group) {
	for i, f := range v.methodOutputs {
		// names := v.vars.Outputs[i]
		// if len(f.Names) == 0 {
		// 	out.Add(
		// 		v.newParamType(p.Type),
		// 	)
		// } else {

		out.
			ListFunc(
				func(out *jen.Group) {
					for _, name := range v.vars.Outputs[i] {
						out.Id(name)
					}
				},
			).
			Add(v.newType(f.Type))
	}
}
