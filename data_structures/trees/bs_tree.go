package trees

type Node struct {
	value               int
	data                interface{}
	left, right, parent *Node
}

type BSTree struct {
	root *Node
}

func NewBSTree(value int, data interface{}) *BSTree {
	return &BSTree{root: &Node{value: value, data: data}}
}

func (tree *BSTree) Insert(value int, data interface{}) *Node {
	return insertNode(tree.root, &Node{value: value, data: data})
}

func (tree *BSTree) Search(value int) (bool, *Node) {
	return searchNode(tree.root, value)
}

func searchNode(start *Node, value int) (bool, *Node) {
	if start == nil {
		return false, nil
	} else if start.value == value {
		return true, start
	} else if value > start.value {
		return searchNode(start.right, value)
	} else {
		return searchNode(start.left, value)
	}
}

func insertNode(start *Node, node *Node) *Node {
	if node.value > start.value {
		if start.right == nil {
			node.parent = start
			start.right = node
		} else {
			return insertNode(start.right, node)
		}
	} else {
		if start.left == nil {
			node.parent = start
			start.left = node
		} else {
			return insertNode(start.left, node)
		}
	}
	return node
}
