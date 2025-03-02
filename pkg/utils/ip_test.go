package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/thealonemusk/WarpNet/pkg/utils"
)

var _ = Describe("IP", func() {
	Context("NextIP", func() {
		It("gives a new IP", func() {
			Expect(NextIP("10.1.1.0", []string{"1.1.0.1"})).To(Equal("1.1.0.2"))
		})
		It("return default", func() {
			Expect(NextIP("10.1.1.0", []string{})).To(Equal("10.1.1.0"))
		})
	})
})
