package generator_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/jmalloc/sham/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	_ "github.com/jmalloc/sham/generator/testdata/inputs"  // force test-cache stale if inputs change
	_ "github.com/jmalloc/sham/generator/testdata/outputs" // force test-cache to be stale if expected outputs change.
)

var _ = Describe("func Generate()", func() {
	files, err := ioutil.ReadDir("testdata/inputs")
	if err != nil {
		panic(err)
	}

	var entries []TableEntry

	for _, f := range files {
		fn := f.Name()

		if fn[0] == '_' {
			continue
		}

		entries = append(
			entries,
			Entry(
				strings.TrimSuffix(fn, ".in"),
				path.Join("testdata/inputs/", fn),
				path.Join("testdata/outputs/", fn),
			),
		)
	}

	DescribeTable(
		"it produces the correct output",
		func(src, expect string) {
			w, err := ioutil.TempFile("", "sham")
			Expect(err).ShouldNot(HaveOccurred())
			defer os.Remove(w.Name())

			err = generator.Generate(src, "outputs", w)
			Expect(err).ShouldNot(HaveOccurred())

			var diff bytes.Buffer
			cmd := exec.Command("diff", "-u", expect, w.Name())
			cmd.Stdout = &diff
			cmd.Stderr = &diff

			err = cmd.Run()

			if diff.Len() > 0 {
				Fail(diff.String())
			}

			Expect(err).ShouldNot(HaveOccurred())
		},
		entries...,
	)
})
