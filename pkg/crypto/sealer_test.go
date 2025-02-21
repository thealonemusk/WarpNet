package crypto_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thealonemusk/WarpNet/pkg/utils"

	. "github.com/thealonemusk/WarpNet/pkg/crypto"
)

var _ = Describe("Crypto utilities", func() {
	Context("AESSealer", func() {
		It("Encode/decode", func() {
			key := RandStringRunes(32)
			message := "foo"

			s := &AESSealer{}

			encoded, err := s.Seal(message, key)
			Expect(err).ToNot(HaveOccurred())
			Expect(encoded).ToNot(Equal(key))
			Expect(len(encoded)).To(Equal(62))

			// Encode again
			encoded2, err := s.Seal(message, key)
			Expect(err).ToNot(HaveOccurred())

			// should differ
			Expect(encoded2).ToNot(Equal(encoded))

			// Decrypt and check
			decoded, err := s.Unseal(encoded, key)
			Expect(err).ToNot(HaveOccurred())
			Expect(decoded).To(Equal(message))

			decoded, err = s.Unseal(encoded2, key)
			Expect(err).ToNot(HaveOccurred())
			Expect(decoded).To(Equal(message))
		})
	})
})
