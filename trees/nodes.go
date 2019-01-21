package trees


type Equaler interface {
	Equal(Equaler) (comparable bool, equal bool)
}


type comparison int8


const (
	LESSER = comparison(-1)
	GREATER = comparison(1)
	EQUAL = comparison(0)
	INCOMPARABLE = comparison(-2)
)


type Comparer interface {
	Equaler
	Compare(Comparer) (result comparison)
}


type Int64 int64


func (i Int64) Equal(j Equaler) (comparable bool, equal bool) {
	if val, ok := j.(Int64); !ok {
		return false, false
	} else {
		result := i.compare(val)
		return result != INCOMPARABLE, result == EQUAL
	}
}


func (i Int64) compare(j Int64) (result comparison) {
	var iVal, jVal = int64(i), int64(j)

	if iVal > jVal {
		return GREATER
	} else if jVal > iVal {
		return LESSER
	}

	return EQUAL
}


func (i Int64) Compare(j Comparer) (result comparison) {
	val, ok := j.(Int64)
	if !ok {
		return INCOMPARABLE
	}

	return i.compare(val)
}


type String string


func (s String) Equal(t Equaler) (comparable bool, equal bool) {
	if val, ok := t.(String); !ok {
		return false, false
	} else {
		result := s.compare(val)
		return result != INCOMPARABLE, result == EQUAL
	}
}

func (s String) compare(t String) (result comparison) {
	var sVal, tVal = string(s), string(t)

	if sVal < tVal {
		return LESSER
	} else if tVal < sVal {
		return GREATER
	}

	return EQUAL
}


func (s String) Compare(t Comparer) (result comparison) {
	val, ok := t.(String)
	if !ok {
		return INCOMPARABLE
	}

	return s.compare(val)
}


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
