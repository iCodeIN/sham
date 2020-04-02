package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

// Generate produces fakes of interfaces parsed from r.
func Generate(src, pkg string, w io.Writer) error {
	v := &visitor{
		OutPackageName: pkg,
	}

	src, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	v.InPackagePath, err = packagePath(src)
	if err != nil {
		return err
	}

	var fs token.FileSet

	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	root, err := parser.ParseFile(
		&fs,
		src,
		data,
		parser.AllErrors,
	)
	if err != nil {
		return err
	}

	ast.Walk(v, root)

	if v.Err != nil {
		return v.Err
	}

	_, err = fmt.Fprintf(w, "%#v", v.Out)
	return err
}

type visitor struct {
	InPackagePath  string
	InPackageName  string
	OutPackageName string
	Out            *jen.File
	Err            error
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil || v.Err != nil {
		return nil
	}

	var descend = true

	switch n := n.(type) {
	case *ast.File:
		descend, v.Err = v.visitFile(n)
	case *ast.TypeSpec:
		descend, v.Err = v.visitTypeSpec(n)
	}

	if descend && v.Err == nil {
		return v
	}

	return nil
}

func (v *visitor) visitFile(n *ast.File) (bool, error) {
	v.InPackageName = n.Name.Name

	v.Out = jen.NewFile(v.OutPackageName)
	v.Out.HeaderComment("Code generated by sham. DO NOT EDIT.")

	return true, nil
}

func (v *visitor) visitTypeSpec(spec *ast.TypeSpec) (bool, error) {
	if !spec.Name.IsExported() {
		return false, nil
	}

	_, ok := spec.Type.(*ast.InterfaceType)
	if !ok {
		return false, nil
	}

	v.Out.Commentf(
		"%s is a test implementation of the %s.%s interface.",
		spec.Name.Name,
		v.InPackageName,
		spec.Name.Name,
	)

	v.Out.
		Type().
		Id(spec.Name.Name).
		StructFunc(func(grp *jen.Group) {
			// for _, m := range iface.Methods.List {
			// 	name := m.Names[0]
			// 	if !name.IsExported() {
			// 		continue
			// 	}

			// 	grp.Id(stubName(m)).
			// 		Func().
			// 		ParamsFunc(func(grp *jen.Group) {
			// 		})
			// }
		})

	v.Out.
		Commentf(
			"Assert that %s implements %s.%s.",
			spec.Name.Name,
			v.InPackageName,
			spec.Name.Name,
		)

	v.Out.
		Var().
		Id("_").
		Qual(v.InPackagePath, spec.Name.Name).
		Op("=").
		Parens(
			jen.Id("*" + spec.Name.Name),
		).
		Parens(
			jen.Nil(),
		)

	// for _, m := range iface.Methods.List {
	// 	name := m.Names[0]
	// 	if !name.IsExported() {
	// 		continue
	// 	}

	// 	v.Out.
	// 		Func().
	// 		Params(
	// 			jen.Id("x").
	// 				Id("*" + spec.Name.Name),
	// 		).
	// 		Id(name.Name).
	// 		ParamsFunc(func(grp *jen.Group) {
	// 		}).
	// 		BlockFunc(func(grp *jen.Group) {
	// 			grp.
	// 				If(
	// 					jen.Id("x").Dot(stubName(m)).Op("!=").Nil(),
	// 				).
	// 				Block(
	// 					jen.Id("x").Dot(stubName(m)).Call(),
	// 				)
	// 		})
	// }

	return false, nil
}

func stubName(n *ast.Field) string {
	return n.Names[0].Name + "Stub"
}