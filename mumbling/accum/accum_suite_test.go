package accum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAccum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Accum Suite")
}
