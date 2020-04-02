package generator

import (
	"fmt"
	"go/ast"
)

// nameTable keeps track of the variable names used within a method.
type nameTable struct {
	used map[string]struct{}
}

func newNameTable(t *ast.FuncType) *nameTable {
	nt := &nameTable{
		used: map[string]struct{}{},
	}

	if t.Params != nil {
		for _, p := range t.Params.List {
			for _, n := range p.Names {
				nt.used[n.Name] = struct{}{}
			}
		}
	}

	if t.Results != nil {
		for _, p := range t.Results.List {
			for _, n := range p.Names {
				nt.used[n.Name] = struct{}{}
			}
		}
	}

	return nt
}

// Get obtains a variable name.
func (nt *nameTable) Get(prefix string) string {
	count := 0

	for {
		candidate := fmt.Sprintf("%s%d", prefix, count)

		if _, ok := nt.used[candidate]; !ok {
			nt.used[candidate] = struct{}{}
			return candidate
		}

		count++
	}
}
