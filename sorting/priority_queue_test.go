package sorting


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sonaak/nettlebane/interfaces"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Priority Queue")
}


var _ = Describe("PriorityQueues", func() {
	Context("Maximum", func(){
		It("should error out when queue is empty", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, err := queue.Maximum()
			Expect(err).ToNot(BeNil())
		})

		It("should return the only element when queue contains only one element", func() {
			heap := NewMaxHeap([]Comparer{
				Int64(12),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			max, err := queue.Maximum()

			Expect(err).To(BeNil())
			Expect(max).To(Equal(Int64(12)))
		})

		It("should return the maximum element of a queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(10),
				Int64(18),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			max, err := queue.Maximum()

			Expect(err).To(BeNil())
			Expect(max).To(Equal(Int64(18)))
		})
	})

	Context("Peek", func() {
		It("should report empty when the queue is empty", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, isEmpty := queue.Peek()
			Expect(isEmpty).To(BeTrue())
		})

		It("should return objects of type Comparer", func() {
			heap := NewMaxHeap([]Comparer{
				Int64(12),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, isEmpty := queue.Peek()

			Expect(isEmpty).To(BeFalse())
			Expect(next).To(Equal(Int64(12)))
		})

		It("should provide the maximum of the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(10),
				Int64(18),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, _ := queue.Peek()
			max, _ := queue.Maximum()
			Expect(next).To(Equal(max))
		})

		It("should not remove any items from the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(10),
				Int64(18),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, _ := queue.Peek()
			Expect(next).To(Equal(Int64(18)))

			Expect(queue.MaxHeap).To(ConsistOf(MaxHeap([]Comparer{
				Int64(18),
				Int64(18),
				Int64(12),
				Int64(10),
			})))
		})

	})

	Context("PopMaximum", func() {
		It("should report error when the queue is empty", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, err := queue.PopMaximum()
			Expect(err).ToNot(BeNil())
			Expect(queue.MaxHeap).To(BeEmpty())
		})

		It("should provide the maximum of the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(10),
				Int64(18),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, _ := queue.Peek()
			max, _ := queue.PopMaximum()
			Expect(next).To(Equal(max))
		})

		It("should remove an item from the front of the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(18),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, _ = queue.PopMaximum()
			Expect(queue.MaxHeap).To(ConsistOf(
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			))
		})

		It("should preserve the heap property for queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(18),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			})

			queue := PriorityQueue{*heap}
			_, _ = queue.PopMaximum()
			Expect(IsMaxHeap([]Comparer(queue.MaxHeap))).To(BeTrue())

		})
	})

	Context("Pop", func() {
		It("should report empty when the queue is empty", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, isEmpty := queue.Pop()
			Expect(isEmpty).To(BeTrue())
		})

		It("should return objects of type Comparer", func() {
			heap := NewMaxHeap([]Comparer{
				Int64(12),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, isEmpty := queue.Pop()

			Expect(isEmpty).To(BeFalse())
			Expect(next).To(Equal(Int64(12)))
		})

		It("should provide the maximum of the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(10),
				Int64(18),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			next, _ := queue.Peek()
			max, _ := queue.Pop()
			Expect(next).To(Equal(max))
		})

		It("should remove an item from the front of the queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(18),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, _ = queue.Pop()
			Expect(queue.MaxHeap).To(ConsistOf(
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			))
		})

		It("should preserve the heap property for queue", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(18),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			})

			queue := PriorityQueue{*heap}
			_, _ = queue.Pop()
			Expect(IsMaxHeap([]Comparer(queue.MaxHeap))).To(BeTrue())
		})
	})

	Context("PushComparer", func() {
		It("should be able to add one object", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			err := queue.PushComparer(Int64(3))
			Expect(err).To(BeNil())
			Expect(queue.MaxHeap).To(Equal(MaxHeap([]Comparer{
				Int64(3),
			})))
		})

		It("should correctly add one object", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(2),
				Int64(10),
				Int64(8),
				Int64(18),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			err := queue.PushComparer(Int64(4))
			Expect(err).To(BeNil())

			Expect(queue.MaxHeap).To(ConsistOf(
				Int64(2),
				Int64(10),
				Int64(8),
				Int64(18),
				Int64(4),
			))
		})

		It("should maintain the heap property when adding objects", func() {
			heap := NewMaxHeap([]Comparer {
				Int64(12),
				Int64(-3),
				Int64(10),
				Int64(3),
				Int64(18),
				Int64(5),
				Int64(18),
				Int64(-1),
				Int64(12),
				Int64(8),
				Int64(0),
				Int64(1),
			})

			queue := PriorityQueue{*heap}
			err := queue.PushComparer(Int64(9))
			Expect(err).To(BeNil())
			Expect(IsMaxHeap([]Comparer(queue.MaxHeap))).To(BeTrue())
		})

		It("should bubble up maximum objects when adding", func() {

		})

		It("should raise error when adding incomtablish objects", func() {

		})
	})
})
