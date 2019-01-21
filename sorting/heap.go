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
	return (*MaxHeap)(&comparers)
}

func (heap *MaxHeap) size() uint64 {
	// nil heaps are empty
	if heap == nil {
		return 0
	}

	return uint64(len(([]Comparer)(*heap)))
}

func (heap *MaxHeap) order(index uint64) {
	if heap == nil {
		return
	}

	data := ([]Comparer)(*heap)
	leftIndex, rightIndex := int64(left(index)), int64(right(index))
	largest := int64(index)
	maxIndex := int64(heap.size()) - 1
	if largest > maxIndex {
		return
	}

	if leftIndex <= maxIndex && data[largest].Compare(data[leftIndex]) == LESSER {
		largest = leftIndex
	}

	if rightIndex <= maxIndex && data[largest].Compare(data[rightIndex]) == LESSER {
		largest = rightIndex
	}

	if largest != int64(index) {
		data[largest], data[index] = data[index], data[largest]
		heap.order(uint64(largest))
	}
}
