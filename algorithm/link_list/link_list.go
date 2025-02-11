package link_list

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

func NewNode[T any](prev *Node[T], data T) *Node[T] {
	newNode := new(Node[T])

	if prev != nil {
		prev.Next = newNode
	}

	newNode.Value = data

	return newNode
}
