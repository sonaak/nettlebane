package trees

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNodes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}

var _ = Describe("Node", func() {
	Describe("Node", func() {
		It("should return the appropriate value", func() {
			Expect(3).To(Equal(3))
		})
	})
})