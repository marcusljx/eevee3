package expressiontree

type Node[T float64] struct {
	Operation       Operation[T]
	DescriptionFunc func(left string, right string) string
	Left            *Node[T]
	Right           *Node[T]
}

func (n *Node[T]) WalkInOrder(walkFunc func(*Node[T])) {
	if n == nil {
		return
	}
	n.Left.WalkInOrder(walkFunc)
	walkFunc(n)
	n.Right.WalkInOrder(walkFunc)
}

func (n *Node[T]) Value() T {
	return n.Operation(n.Left.Value(), n.Right.Value())
}

func (n *Node[T]) Equals(other *Node[T]) bool {
	if n.Operation.String() != other.Operation.String() {
		return false
	}
	if !n.Left.Equals(other.Left) {
		return false
	}
	if !n.Right.Equals(other.Right) {
		return false
	}
	return true
}

func (n *Node[T]) Describe() string {
	return n.DescriptionFunc(n.Left.Describe(), n.Right.Describe())
}
