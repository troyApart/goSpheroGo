package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoSpheroGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoSpheroGo Suite")
}
