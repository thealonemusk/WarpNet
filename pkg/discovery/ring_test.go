package discovery_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/thealonemusk/WarpNet/pkg/discovery"
)

var _ = Describe("String utilities", func() {
	Context("Ring", func() {
		It("adds elements to the ring (3)", func() {
			R := Ring{Length: 3}
			R.Add("a")
			R.Add("b")
			R.Add("c")
			Expect(R.Data).To(Equal([]string{"a", "b", "c"}))
			R.Add("d")
			Expect(R.Data).To(Equal([]string{"b", "c", "d"}))
			R.Add("d")
			Expect(R.Data).To(Equal([]string{"b", "c", "d"}))
		})
		It("adds elements to the ring (2)", func() {
			R := Ring{Length: 2}
			R.Add("a")
			R.Add("b")
			R.Add("c")
			Expect(R.Data).To(Equal([]string{"b", "c"}))
			R.Add("d")
			Expect(R.Data).To(Equal([]string{"c", "d"}))
			R.Add("d")
			Expect(R.Data).To(Equal([]string{"c", "d"}))
		})
	})
})
