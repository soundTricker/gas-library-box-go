package gaslibrarybox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGaslibrarybox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gaslibrarybox Suite")
}
