package service_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/thealonemusk/WarpNet/api/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testInstance = os.Getenv("TEST_INSTANCE")

func TestService(t *testing.T) {
	if testInstance == "" {
		fmt.Println("a testing instance has to be defined with TEST_INSTANCE")
		os.Exit(1)
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	// Start the test suite only if we have some machines connected

	Eventually(func() (int, error) {
		c := NewClient(WithHost(testInstance))
		m, err := c.Machines()
		return len(m), err
	}, 100*time.Second, 1*time.Second).Should(BeNumerically(">=", 0))
})
