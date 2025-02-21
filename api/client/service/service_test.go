package service_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	client "github.com/thealonemusk/WarpNet/api/client"

	. "github.com/thealonemusk/WarpNet/api/client/service"
)

var _ = Describe("Service", func() {
	c := client.NewClient(client.WithHost(testInstance))
	s := NewClient("foo", c)
	Context("Retrieves nodes", func() {
		PIt("Detect nodes", func() {
			Eventually(func() []string {
				n, _ := s.ActiveNodes()
				return n
			},
				100*time.Second, 1*time.Second).ShouldNot(BeEmpty())
		})
	})

	Context("Advertize nodes", func() {
		It("Detect nodes", func() {
			n, err := s.AdvertizingNodes()
			Expect(len(n)).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())

			s.Advertize("foo")

			Eventually(func() []string {
				n, _ := s.AdvertizingNodes()
				return n
			},
				100*time.Second, 1*time.Second).Should(Equal([]string{"foo"}))
		})
	})
})
