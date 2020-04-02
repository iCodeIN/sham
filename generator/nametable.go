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

// GetNumbered obtains a variable name that always has a numbered suffix.
func (nt *nameTable) GetNumbered(prefix string) string {
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

// Get obtains a variable name, only adding a numbered suffix if there is a
// collision.
func (nt *nameTable) Get(prefix string) string {
	count := 0
	candidate := prefix

	for {

		if _, ok := nt.used[candidate]; !ok {
			nt.used[candidate] = struct{}{}
			return candidate
		}

		candidate = fmt.Sprintf("%s%d", prefix, count)
		count++
	}
}
