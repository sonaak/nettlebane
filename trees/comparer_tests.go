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

		It("should be less another Int64 of greater value", func() {
			result := Int64(10).Compare(Int64(12))
			Expect(result).To(Equal(LESSER))
		})

		It("should be less another Int64 of greater value", func() {
			result := Int64(13).Compare(Int64(12))
			Expect(result).To(Equal(GREATER))
		})
	})
})
