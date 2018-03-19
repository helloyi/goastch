package goastcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoastcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goastcher Suite")
}
