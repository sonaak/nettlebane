package interfaces


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestComparables(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comparables")
}

var _ = Describe("Comparers", func() {
	Describe("Int64", func() {
		It("should equal another Int64 of same value", func() {
			comparable, same := Int64(12).Equal(Int64(12))
			Expect(comparable).To(BeTrue())
			Expect(same).To(BeTrue())
		})


		It("should not equal another Int64 of different value", func() {
			comparable, same := Int64(12).Equal(Int64(10))
			Expect(comparable).To(BeTrue())
			Expect(same).To(BeFalse())
		})

		It("should not be comparable to another Equaler of different type", func() {
			comparable, _ := Int64(12).Equal(String("12"))
			Expect(comparable).To(BeFalse())
		})

		It("should be less another Int64 of greater value", func() {
			result := Int64(10).Compare(Int64(12))
			Expect(result).To(Equal(LESSER))
		})

		It("should be less another Int64 of greater value", func() {
			result := Int64(13).Compare(Int64(12))
			Expect(result).To(Equal(GREATER))
		})

		It("should not be comparable to another Comparer of different type", func() {
			result := Int64(13).Compare(String("13"))
			Expect(result).To(Equal(INCOMPARABLE))
		})
	})

	Describe("String", func(){
		It("should equal another String of same value", func() {
			comparable, same := String("hello").Equal(String("hello"))
			Expect(comparable).To(BeTrue())
			Expect(same).To(BeTrue())
		})

		It("should not another String of different value", func() {
			comparable, same := String("hello").Equal(String("hello "))
			Expect(comparable).To(BeTrue())
			Expect(same).To(BeFalse())
		})

		It("should not be comparable to another Equaler of different type", func() {
			comparable, _ := String("12").Equal(Int64(12))
			Expect(comparable).To(BeFalse())
		})

		It("should be less than another String of greater value", func() {
			result := String("hello").Compare(String("world"))
			Expect(result).To(Equal(LESSER))
		})

		It("should be greater than another String of lesser value", func() {
			result := String("hello").Compare(String("Hello"))
			Expect(result).To(Equal(GREATER))
		})

		It("should not be comparable to another Comparer of different type", func() {
			result := String("13").Compare(Int64(13))
			Expect(result).To(Equal(INCOMPARABLE))
		})
	})
})
