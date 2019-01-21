package trees

import (
	. "github.com/sonaak/nettlebane/interfaces"
)

type Node struct {
	value Equaler
	childNodes []*Node
}


func NewNode(value Equaler, childNodes []*Node) *Node {
	return &Node {
		value: value,
		childNodes: childNodes,
	}
}


func (node *Node) Value() Equaler {
	return node.value
}


func (node *Node) Children() []*Node {
	return node.childNodes
}


func (node *Node) AppendChild(childNode *Node) {
	node.childNodes = append(node.childNodes, childNode)
}


func (node *Node) PrependChild(childNode *Node) {
	node.childNodes = append([]*Node{childNode}, node.childNodes...)
}


type BinaryNode Node


func NewBinaryNode(value Comparer, left, right *BinaryNode) *BinaryNode {
	node := Node {
		value: value,
		childNodes: []*Node {
			(*Node)(left), (*Node)(right),
		},
	}

	return (*BinaryNode)(&node)
}


func (node *BinaryNode) Comparer() Comparer {
	comp, _ := node.value.(Comparer)
	return comp
}


func (node *BinaryNode) Left() *BinaryNode {
	leftNode := node.childNodes[0]
	return (*BinaryNode)(leftNode)
}


func (node *BinaryNode) Right() *BinaryNode {
	leftNode := node.childNodes[1]
	return (*BinaryNode)(leftNode)
}

// Assume that left is nil, sets left to n
func (node *BinaryNode) SetLeft(n *BinaryNode) {
	node.childNodes[0] = (*Node)(n)
}

// Assume that right is nil, sets right to n
func (node *BinaryNode) SetRight(n *BinaryNode) {
	node.childNodes[1] = (*Node)(n)
}
