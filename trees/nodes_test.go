package trees

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNodes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}

var _ = Describe("Node", func() {
	Describe("Node", func() {
		Context("NewNode", func(){
			It("should return the right value and children", func() {
				nodes := make([]*Node, 0)
				value := String("hello")
				node := NewNode(value, nodes)
				Expect(node.value).To(Equal(value))
				Expect(node.childNodes).To(Equal(nodes))
			})
		})
		Context("Value", func() {
			It("should return the correct value", func(){
				nodes := make([]*Node, 0)
				value := String("hello")
				node := NewNode(value, nodes)
				Expect(node.Value()).To(Equal(value))
			})
		})

		Context("Children", func() {
			It("should return the correct Children", func(){
				nodes := make([]*Node, 0)
				value := String("hello")
				node := NewNode(value, nodes)
				Expect(node.Children()).To(Equal(nodes))
			})
		})
	})

	Describe("BinaryNode", func() {
		Context("NewBinaryNode", func() {
			It("should construct the right value and children", func(){
				leftNode, rightNode := BinaryNode(Node{}), BinaryNode(Node{})
				node := NewBinaryNode(String("hello"), &leftNode, &rightNode)

				Expect((*Node)(node).Value()).Should(Equal(String("hello")))
				Expect((*Node)(node).Children()).Should(ConsistOf(
					(*Node)(&leftNode),
					(*Node)(&rightNode),
				))
			})
		})

		Context("Comparer", func() {
			It("should return the right value", func(){
				leftNode, rightNode := BinaryNode(Node{}), BinaryNode(Node{})
				node := NewBinaryNode(String("hello"), &leftNode, &rightNode)
				Expect(node.Comparer()).Should(Equal(String("hello")))
			})
		})

		Context("Left", func() {
			It("should return the correct left node", func(){
				leftNode, rightNode := BinaryNode(Node{}), BinaryNode(Node{})
				node := NewBinaryNode(String("hello"), &leftNode, &rightNode)
				Expect(node.Left()).Should(Equal(&leftNode))
			})
		})

		Context("Right", func() {
			It("should return the correct right node", func(){
				leftNode, rightNode := BinaryNode(Node{}), BinaryNode(Node{})
				node := NewBinaryNode(String("hello"), &leftNode, &rightNode)
				Expect(node.Right()).Should(Equal(&rightNode))
			})
		})
	})
})