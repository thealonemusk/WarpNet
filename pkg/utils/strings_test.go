package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/thealonemusk/WarpNet/pkg/utils"
)

var _ = Describe("String utilities", func() {
	Context("RandStringRunes", func() {
		It("returns a string with the correct length", func() {
			Expect(len(RandStringRunes(10))).To(Equal(10))
		})
	})
})
