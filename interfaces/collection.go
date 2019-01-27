package interfaces


type PushPopper interface {
	// Looks at the next item without modifying the collection
	Peek() (interface{}, bool)

	// Push an object into the collection; modifies the collection
	Push(interface{}) error

	// Pops an object from the collection; modifies the collection
	Pop() (interface{}, bool)
}
