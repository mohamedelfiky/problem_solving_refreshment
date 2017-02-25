package linked_list

type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) Push(value interface{}) *ListNode {
	node := &ListNode{value: value}
	if l.head == nil {
		l.head = node
	} else {
		last := &(l.head)
		for (*last).next != nil {
			last = &((*last).next)
		}
		(*last).next = node
	}
	return node
}

func (l *LinkedList) Insert(afterMe *ListNode, value interface{}) *ListNode {
	node := &ListNode{value: value}
	node.next = afterMe.next
	afterMe.next = node
	return node
}
