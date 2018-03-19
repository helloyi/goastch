package scanner_test

import (
	. "github.com/helloyi/goastch/galang/scanner"
	. "github.com/helloyi/goastch/galang/token"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	"fmt"
	"strings"
)

var _ = Describe("Scanner", func() {
	Describe("binding", func() {
		s := New(strings.NewReader("structType has fieldList has @field isType \"test\""))
		It("should no error", func() {
			for tok := s.Scan(); tok != EOF; tok = s.Scan() {
				fmt.Printf("%s: %s\n", s.Position, s.TokenText())
				// Expect(skiplist.MaxLevel()).To(Equal(DefaultMaxLevel))
				// Expect(skiplist.Probability()).To(Equal(DefaultProbability))
			}
		})
	})
})
