package goastch_test

import (
	"bytes"
	"go/build"
	"go/parser"
	"go/printer"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"golang.org/x/tools/go/loader"

	"github.com/helloyi/goastch"
	. "github.com/helloyi/goastch/goastcher"
)

var _ = ginkgo.Describe("Goastch", func() {
	var (
		ger   = HasDescendant(File(Has(ImportSpec(HasName(Equals("."))).Bind(""))))
		fname = "<buffer>"
		code  = `
// Package foo
package foo

import (
	"fmt"
	. "golang.org/x/net/context"
)

type aaa struct {
	ctx Context
}

func bar() bool {
	xx := aaa{}
	fmt.Println(xx)
	return false
}`
	)
	ginkgo.It("test", func() {
		loader := &loader.Config{
			Build:      &build.Default,
			ParserMode: parser.ParseComments,
		}
		file, err := loader.ParseFile(fname, code)
		gomega.Expect(err).Should(gomega.BeNil())

		loader.CreateFromFiles(fname, file)
		prog, err := loader.Load()
		gomega.Expect(err).Should(gomega.BeNil())

		pkgInfo := prog.Package(fname)
		info := &pkgInfo.Info
		fset := loader.Fset

		result, err := goastch.Find(file, info, ger)
		gomega.Expect(err).Should(gomega.BeNil())

		buf := new(bytes.Buffer)
		gomega.Expect(len(result)).Should(gomega.Equal(1))
		for _, list := range result {
			gomega.Expect(len(list)).Should(gomega.Equal(1))
			for _, node := range list {
				err = printer.Fprint(buf, fset, node)
				gomega.Expect(err).Should(gomega.BeNil())
				gomega.Expect(buf.String()).Should(gomega.Equal(`. "golang.org/x/net/context"`))
			}
		}
	})
})
