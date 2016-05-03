package snow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSnow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snow Suite")
}
