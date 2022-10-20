package go_stack_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoStack Suite")
}
