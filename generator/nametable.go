package generator

import "fmt"

// nameTable keeps track of the variable names used within a scope.
type nameTable struct {
	prefix  string
	used    map[string]struct{}
	counter int
}

// Get obtains a variable name.
func (nt *nameTable) Get(desired string) string {
	if desired != "" {
		return nt.get(desired)
	}

	nt.counter++

	return nt.get(
		fmt.Sprintf(
			"%s%d",
			nt.prefix,
			nt.counter-1,
		),
	)
}

func (nt *nameTable) get(n string) string {
	count := 0
	candidate := n

	for {
		if _, ok := nt.used[candidate]; !ok {
			return candidate
		}

		candidate = fmt.Sprintf("%s%d", n, count)
		count++
	}
}
