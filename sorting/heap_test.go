package sorting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sonaak/nettlebane/interfaces"
	"testing"
)

func TestHeaps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}

var _ = Describe("MaxHeaps", func() {
	Context("NewMaxHeaps", func() {
		It("should instantiate an empty list", func() {
			comparers := []Comparer{}
			Expect(func() {
				NewMaxHeap(comparers)
			}).ToNot(Panic())
		})
	})

	Context("order", func() {
		It("should order correctly", func() {
		 	intList := []Comparer {
		 		Int64(16),
		 		Int64(4),
		 		Int64(10),
		 		Int64(14),
		 		Int64(7),
		 		Int64(9),
		 		Int64(3),
		 		Int64(2),
		 		Int64(8),
		 		Int64(1),
			}
			heap := MaxHeap(intList)
			heap.order(1)

			Expect(heap).To(Equal(MaxHeap([]Comparer {
				Int64(16),
		 		Int64(14),
		 		Int64(10),
		 		Int64(8),
		 		Int64(7),
		 		Int64(9),
		 		Int64(3),
		 		Int64(2),
		 		Int64(4),
		 		Int64(1),
			})))
		})

		It("should stop when already ordered", func() {
			intList := []Comparer {
		 		Int64(16),
		 		Int64(14),
		 		Int64(10),
		 		Int64(4),
		 		Int64(7),
		 		Int64(9),
		 		Int64(3),
		 		Int64(2),
		 		Int64(8),
		 		Int64(1),
			}
			heap := MaxHeap(intList)
			heap.order(1)

			Expect(heap).To(Equal(MaxHeap([]Comparer {
				Int64(16),
		 		Int64(14),
		 		Int64(10),
		 		Int64(4),
		 		Int64(7),
		 		Int64(9),
		 		Int64(3),
		 		Int64(2),
		 		Int64(8),
		 		Int64(1),
			})))
		})

		It("should do fine with index out of bounds", func() {
			intList := []Comparer{}
			heap := MaxHeap(intList)
			Expect(func() {heap.order(2)}).ToNot(Panic())
		})
	})
})
