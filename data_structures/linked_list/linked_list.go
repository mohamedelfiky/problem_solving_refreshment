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

func (l *LinkedList) Search(value interface{}) *ListNode {
	current := &(l.head)
	for *current != nil {
		if (*current).value == value {
			return *current
		}
		current = &((*current).next)
	}
	return nil
}

func (l *LinkedList) PushFirst(value interface{}) {
	l.head = &ListNode{value: value, next: l.head}
}

func (l *LinkedList) Remove(item *ListNode) {
	if l.head != nil {
		current := &(l.head)
		for *current != nil {
			if (*current).next == item {
				(*current).next = (*current).next.next
				break
			}
			current = &((*current).next)
		}
	}
}
