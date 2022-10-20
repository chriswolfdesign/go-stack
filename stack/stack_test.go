package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	stackLib "go-stack/stack"
)

var _ = Describe("Stack", func() {
	var stack stackLib.Stack

	BeforeEach(func() {
		stack = stackLib.Stack{}
	})

	Describe("Initialization tests", func() {
		When("Stack is initialized", func() {
			It("Has a size of 0", func() {
				Expect(stack.Size()).To(Equal(0))
			})
		})
	})

	Describe("Push tests", func() {
		When("pushing to empty stack", func() {
			It("has size 1", func() {
				stack.Push(1)
				Expect(stack.Size()).To(Equal(1))
			})
		})

		When("pushing to stack with one element", func() {
			BeforeEach(func() {
				stack.Push(1)
			})

			It("has size 1", func() {
				stack.Push(2)
				Expect(stack.Size()).To(Equal(2))
			})
		})

		When("pushing to stack with multiple elements", func() {
			BeforeEach(func() {
				stack.Push(1)
				stack.Push(2)
			})

			It("has size 2", func() {
				stack.Push(3)
				Expect(stack.Size()).To(Equal(3))
			})
		})
	})
})
