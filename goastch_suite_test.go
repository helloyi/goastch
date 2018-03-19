package goastch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoastch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goastch Suite")
}
