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
				tail, _ := stack.Peek()
				Expect(tail).To(Equal(1))
			})
		})

		When("pushing to stack with one element", func() {
			BeforeEach(func() {
				stack.Push(1)
			})

			It("has size 1", func() {
				stack.Push(2)
				Expect(stack.Size()).To(Equal(2))
				tail, _ := stack.Peek()
				Expect(tail).To(Equal(2))
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
				tail, _ := stack.Peek()
				Expect(tail).To(Equal(3))
			})
		})
	})

	Describe("Peek tests", func() {
		When("stack is empty", func() {
			It("should throw an error", func() {
				head, err := stack.Peek()
				Expect(err).To(HaveOccurred())
				Expect(head).To(Equal(0))
			})
		})

		When("stack only has on element", func() {
			BeforeEach(func() {
				stack.Push(1)
			})

			It("returns that value and does not throw an error", func() {
				tail, err := stack.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(tail).To(Equal(1))
			})
		})

		When("stack has multiple elements", func() {
			BeforeEach(func() {
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
			})

			It("returns the last element in the stack and does not throw an error", func() {
				tail, err := stack.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(tail).To(Equal(3))
			})
		})
	})

	Describe("Pop tests", func() {
		When("popping from empty stack", func() {
			It("throws an error", func() {
				element, err := stack.Pop()
				Expect(err).To(HaveOccurred())
				Expect(element).To(Equal(0))
			})
		})

		When("popping from stack with one element", func() {
			BeforeEach(func() {
				stack.Push(1)
			})

			It("removes that element and does not throw an error", func() {
				element, err := stack.Pop()
				Expect(err).ToNot(HaveOccurred())
				Expect(element).To(Equal(1))
				tail, err := stack.Peek()
				Expect(err).To(HaveOccurred())
				Expect(tail).To(Equal(0))
			})
		})

		When("popping from stack with multiple elements", func() {
			BeforeEach(func() {
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
			})
			It("removes last element and does not throw error", func() {
				element, err := stack.Pop()
				Expect(err).ToNot(HaveOccurred())
				Expect(element).To(Equal(3))
				tail, err := stack.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(tail).To(Equal(2))
			})
		})
	})
})
