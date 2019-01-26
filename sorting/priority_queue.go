package sorting

import (
	"github.com/pkg/errors"
	. "github.com/sonaak/nettlebane/interfaces"
)


type PriorityQueue struct {
	MaxHeap
}

func (queue *PriorityQueue) Maxium() (Comparer, error) {
	comparers := ([]Comparer)(queue.MaxHeap)
	if len(comparers) <= 0 {
		return nil, errors.New("priority queue empty")
	}

	return comparers[0], nil
}

func (queue *PriorityQueue) Peek() (interface{}, bool) {
	next, err := queue.Maxium()
	if err != nil {
		return nil, true
	}

	return next, false
}
