package trees


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestComparables(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}

var _ = Describe("Comparers", func() {
	Describe("Int64", func() {
		It("should equal another Int64 of same value", func() {
			comparable, same := Int64(12).Equal(Int64(12))
			Expect(comparable).To(BeTrue())
			Expect(same).To(BeTrue())
		})
	})
})
