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
)

var _ = Describe("func Generate()", func() {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		panic(err)
	}

	var entries []TableEntry

	for _, f := range files {
		fn := f.Name()

		if fn[0] == '_' {
			continue
		}

		if !strings.HasSuffix(fn, ".in") {
			continue
		}

		name := strings.TrimSuffix(fn, ".in")

		entries = append(
			entries,
			Entry(
				name,
				path.Join("testdata", fn),
				path.Join("testdata", name+".out"),
			),
		)
	}

	DescribeTable(
		"it produces the correct output",
		func(input, output string) {
			in, err := os.Open(input)
			Expect(err).ShouldNot(HaveOccurred())

			out, err := ioutil.TempFile("", "sham")
			Expect(err).ShouldNot(HaveOccurred())
			defer os.Remove(out.Name())

			err = generator.Generate(in, out)
			Expect(err).ShouldNot(HaveOccurred())

			var diff bytes.Buffer
			cmd := exec.Command("diff", "-u", output, out.Name())
			cmd.Stdout = &diff

			err = cmd.Run()

			if diff.Len() > 0 {
				Fail(diff.String())
			}

			Expect(err).ShouldNot(HaveOccurred())
		},
		entries...,
	)
})
