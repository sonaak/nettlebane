package sorting


import (
	. "github.com/sonaak/nettlebane/interfaces"
)


func left(index uint64) uint64 {
	return index << 1 | 1
}

func right(index uint64) uint64 {
	return (index + 1) << 1
}

func parent(index uint64) uint64 {
	if index == 0 {
		return 0
	}

	return (index - 1) >> 1
}

// A max binary heap
// See Chapter 6 in Cormen et al. 2e (p. 127)
type MaxHeap []Comparer

func NewMaxHeap(comparers []Comparer) *MaxHeap {
	heap := (*MaxHeap)(&comparers)
	for i := len(comparers) << 1 - 1; i >= 0; i-- {
		heap.order(uint64(i))
	}
	return heap
}

func (heap *MaxHeap) size() uint64 {
	// nil heaps are empty
	if heap == nil {
		return 0
	}

	return uint64(len(([]Comparer)(*heap)))
}


func (heap *MaxHeap) compareAndSwap(current uint64) (largest uint64) {
	leftIndex, rightIndex := left(current), right(current)
	largest = current
	maxIndex := heap.size() - 1
	data := ([]Comparer)(*heap)

	if leftIndex <= maxIndex && data[largest].Compare(data[leftIndex]) == LESSER {
		largest = leftIndex
	}

	if rightIndex <= maxIndex && data[largest].Compare(data[rightIndex]) == LESSER {
		largest = rightIndex
	}

	if largest != current {
		data[largest], data[current] = data[current], data[largest]
	}
	return
}


func (heap *MaxHeap) order(index uint64) {
	if heap.size() == uint64(0) {
		return
	}

	largest := heap.compareAndSwap(index)
	if largest != index {
		heap.order(largest)
	}
}

func (heap *MaxHeap) orderIteratively(index uint64) {
	if heap.size() == uint64(0) {
		return
	}

	for index < heap.size() - 1 {
		largest := heap.compareAndSwap(index)
		if largest == index {
			return
		}
		index = largest
	}

}

// Sorting with MaxHeap in place.
func HeapSort(comparers []Comparer) {
	_ = NewMaxHeap(comparers)
	for i := len(comparers) - 1; i > 0; i++ {
		comparers[0], comparers[i] = comparers[i], comparers[0]
		unsorted := comparers[:i]
		(*MaxHeap)(&unsorted).order(0)
	}
}