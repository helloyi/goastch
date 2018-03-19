package parser_test

import (
	. "github.com/helloyi/goastch/galang/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

var _ = Describe("Parser", func() {
	It("parserGer", func() {
		ger, err := ParseGer("HasDes ShortVarDecl Has @compositeLit IsSize 0")
		fmt.Println("ll", ger)
		Expect(err).To(BeNil())
	})
	It("hasname", func() {
		ger, err := ParseGer("HasDes file has @importSpec hasName equals \".\"")
		fmt.Println(ger)
		Expect(err).To(BeNil())
	})
	It("hasname", func() {
		ger, err := ParseGer(`file hasName allof unless hasSuffix "_test" and contains "_"`)
		fmt.Println(ger)
		Expect(err).To(BeNil())
	})
})
