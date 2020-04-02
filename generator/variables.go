package generator

import (
	"fmt"
	"go/ast"
)

// variables keeps track of the variable names used within a method.
type variables struct {
	Receiver string
	Inputs   [][]string
	Outputs  [][]string
}

func newVariables(t *ast.FuncType) *variables {
	v := &variables{}

	used := map[string]struct{}{}

	if t.Params != nil {
		v.Inputs = make([][]string, len(t.Params.List))

		for i, f := range t.Params.List {
			for _, n := range f.Names {
				v.Inputs[i] = append(v.Inputs[i], n.Name)
				used[n.Name] = struct{}{}
			}
		}
	}

	if t.Results != nil {
		v.Outputs = make([][]string, len(t.Results.List))

		for i, f := range t.Results.List {
			for _, n := range f.Names {
				v.Outputs[i] = append(v.Outputs[i], n.Name)
				used[n.Name] = struct{}{}
			}
		}
	}

	for i, names := range v.Inputs {
		if len(names) == 0 {
			v.Inputs[i] = []string{
				generateName(
					used,
					fmt.Sprintf("arg%d", i),
				),
			}
		}
	}

	v.Receiver = generateName(used, "stub")

	return v
}

func generateName(
	used map[string]struct{},
	desired string,
) string {
	candidate := desired

	for {
		if _, ok := used[candidate]; !ok {
			used[candidate] = struct{}{}
			return candidate
		}

		candidate += "_"
	}
}
