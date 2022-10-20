package go_stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoStack", func() {
	When("should fail", func() {
		It("fails again", func() {
			Expect(false).To(BeTrue())
		})
	})
})
