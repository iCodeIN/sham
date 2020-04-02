package main

import (
	"fmt"
	"os"

	"github.com/jmalloc/sham/generator"
)

func main() {
	if err := generator.Generate(os.Args[0], "fixtures", os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
