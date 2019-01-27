package sorting

import (
	"github.com/pkg/errors"
	. "github.com/sonaak/nettlebane/interfaces"
)


type PriorityQueue struct {
	MaxHeap
}

func (queue *PriorityQueue) Maximum() (Comparer, error) {
	comparers := ([]Comparer)(queue.MaxHeap)
	if len(comparers) <= 0 {
		return nil, errors.New("priority queue empty")
	}

	return comparers[0], nil
}

func (queue *PriorityQueue) Peek() (interface{}, bool) {
	next, err := queue.Maximum()
	if err != nil {
		return nil, true
	}

	return next, false
}

func (queue *PriorityQueue) PopMaximum() (max Comparer, err error) {
	max, err = queue.Maximum()
	if err != nil {
		return
	}

	remainder := []Comparer(queue.MaxHeap)[1:]
	queue.MaxHeap = MaxHeap(remainder)
	queue.MaxHeap.order(0)

	return
}

func (queue *PriorityQueue) Pop() (interface{}, bool) {
	next, err := queue.PopMaximum()
	if err != nil {
		return nil, true
	}

	return next, false
}
