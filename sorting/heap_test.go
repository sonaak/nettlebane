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
	Context("parent", func() {
		It("should define the correct number", func() {
			Expect(parent(1)).To(Equal(uint64(0)))
			Expect(parent(0)).To(Equal(uint64(0)))
			Expect(parent(8)).To(Equal(uint64(3)))
		})
	})

	Context("left/right", func() {
		It("should define the correct number for left", func() {
			Expect(left(1)).To(Equal(uint64(3)))
			Expect(left(0)).To(Equal(uint64(1)))
			Expect(left(3)).To(Equal(uint64(7)))
		})

		It("should define the correct number for right", func() {
			Expect(right(1)).To(Equal(uint64(4)))
			Expect(right(0)).To(Equal(uint64(2)))
			Expect(right(3)).To(Equal(uint64(8)))
		})
	})


	Context("NewMaxHeaps", func() {
		It("should instantiate an empty list", func() {
			comparers := []Comparer{}
			Expect(func() {
				NewMaxHeap(comparers)
			}).ToNot(Panic())
		})

		It("should (max) heapify a list of comparers", func() {
			comparers := []Comparer {
				Int64(4),
				Int64(8),
				Int64(9),
				Int64(12),
				Int64(2),
			}
			Expect(NewMaxHeap(comparers)).To(Equal((*MaxHeap)(&[]Comparer{
				Int64(12),
				Int64(8),
				Int64(9),
				Int64(4),
				Int64(2),
			})))
		})
	})

	Context("size", func() {
		It("should return 0 for empty heaps", func() {
			heap := MaxHeap([]Comparer {})
			Expect(heap.size()).To(Equal(uint64(0)))
		})

		It("should return the correct sizes", func() {


			heap := MaxHeap([]Comparer {
				Int64(9),
				Int64(6),
				Int64(1),
			})
			Expect(heap.size()).To(Equal(uint64(3)))
		})

		It("should not panic for nil heaps", func() {
			var heap *MaxHeap = nil
			Expect(func() {heap.size()}).NotTo(Panic())
			Expect(heap.size()).To(Equal(uint64(0)))
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

		It("should do fine with nil heaps", func() {
			var heap *MaxHeap = nil

			Expect(func() {heap.order(1)}).NotTo(Panic())
		})
	})

	Context("orderIteratively", func() {
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
			heap.orderIteratively(1)

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
			heap.orderIteratively(1)

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
			Expect(func() {heap.orderIteratively(2)}).ToNot(Panic())
		})

		It("should do fine with nil heaps", func() {
			var heap *MaxHeap = nil

			Expect(func() {heap.orderIteratively(1)}).NotTo(Panic())
		})
	})

	Context("HeapSort", func() {
		It("should sort the empty list correctly", func() {
			comparers := []Comparer{}
			Expect(func() {HeapSort(comparers)}).ToNot(Panic())
			Expect(comparers).To(Equal(comparers))
		})

		It("should sort a list of integers correctly", func() {
			comparers := []Comparer{
				Int64(12),
				Int64(1),
				Int64(-2),
				Int64(0),
				Int64(4),
				Int64(18),
				Int64(5),
			}

			HeapSort(comparers)
			Expect(comparers).To(Equal([]Comparer {
				Int64(-2),
				Int64(0),
				Int64(1),
				Int64(4),
				Int64(5),
				Int64(12),
				Int64(18),
			}))
		})

		It("should preserve a sorted list of integers", func() {
			comparers := []Comparer {
				Int64(-1),
				Int64(0),
				Int64(1),
				Int64(2),
				Int64(3),
				Int64(4),
				Int64(5),
			}

			HeapSort(comparers)
			Expect(comparers).To(Equal([]Comparer {
				Int64(-1),
				Int64(0),
				Int64(1),
				Int64(2),
				Int64(3),
				Int64(4),
				Int64(5),
			}))
		})
	})

})
