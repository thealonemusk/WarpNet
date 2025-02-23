package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/thealonemusk/WarpNet/pkg/utils"
)

var _ = Describe("Leader utilities", func() {
	Context("Leader", func() {
		It("returns the correct leader", func() {
			Expect(Leader([]string{"a", "b", "c", "d"})).To(Equal("b"))
			Expect(Leader([]string{"a", "b", "c", "d", "e", "f", "G", "bb"})).To(Equal("b"))
			Expect(Leader([]string{"a", "b", "c", "d", "e", "f", "G", "bb", "z", "b1", "b2"})).To(Equal("z"))
			Expect(Leader([]string{"1", "2", "3", "4", "5"})).To(Equal("2"))
			Expect(Leader([]string{"1", "2", "3", "4", "5", "6", "7", "21", "22"})).To(Equal("22"))
		})
	})
})
