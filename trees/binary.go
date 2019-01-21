package trees

import (
	. "github.com/sonaak/nettlebane/interfaces"
	"sync"
)

// A tree that satisfies the "binary search" property, and therefore
// has a unique insert/remove action
type BinarySearchTree struct {
	root *BinaryNode

	lock *sync.RWMutex
}

func (tree *BinarySearchTree) Insert(value Comparer) {
	// find the node and insert
}

func (tree *BinarySearchTree) Remove(value Comparer) (removed bool) {
	//
	return
}

func (tree *BinarySearchTree) Find(value Comparer) (node Node, found bool) {
	return
}

