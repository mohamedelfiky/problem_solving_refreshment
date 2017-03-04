package tree

type Node struct {
	value               int
	data                interface{}
	left, right, parent *Node
}

type Tree struct {
	top Node
}
