package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go-stack/stack"
)

var _ = Describe("Stack", func() {
	var stack stack.Stack

	Describe("Initialization tests", func() {
		When("Stack is initialized", func() {
			It("Has a size of 0", func() {
				Expect(stack.Size).To(Equal(0))
			})
		})
	})

})
