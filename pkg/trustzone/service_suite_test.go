package trustzone_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTrustzone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trustzone Suite")
}
