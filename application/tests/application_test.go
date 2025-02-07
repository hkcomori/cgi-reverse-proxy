package tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApplicationSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing application layer")
}
