package sorting


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sonaak/nettlebane/interfaces"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}


var _ = Describe("PriorityQueues", func() {
	Context("Maximum", func(){
		It("should error out when queue is empty", func() {
			empty := []Comparer{}
			heap := NewMaxHeap(empty)
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			_, err := queue.Maxium()
			Expect(err).ToNot(BeNil())
		})

		It("should return the only element when queue contains only one element", func() {
			heap := NewMaxHeap([]Comparer{
				Int64(12),
			})
			Expect(heap).ToNot(BeNil())

			queue := PriorityQueue{*heap}
			max, err := queue.Maxium()

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
			max, err := queue.Maxium()

			Expect(err).To(BeNil())
			Expect(max).To(Equal(Int64(18)))
		})
	})
})
