package doubly_linked_list

type ListNode struct {
	value      interface{}
	next, prev *ListNode
}

type DoublyLinkedList struct {
	head, tail *ListNode
}

func (l *DoublyLinkedList) Push(value interface{}) *ListNode {
	node := &ListNode{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	return node
}

func (l *DoublyLinkedList) Insert(afterMe *ListNode, value interface{}) *ListNode {
	node := &ListNode{value: value, prev: afterMe, next: afterMe.next}
	afterMe.next = node
	if node.next == nil {
		l.tail = node
	} else {
		node.next.prev = node
	}
	return node
}

func (l *DoublyLinkedList) Search(value interface{}) *ListNode {
	if l.head != nil {
		current := &(l.head)
		for *current != nil {
			if (*current).value == value {
				return *current
			}
			current = &((*current).next)
		}
	}
	return nil
}

func (l *DoublyLinkedList) PushFirst(value interface{}) *ListNode {
	l.head = &ListNode{value: value, next: l.head}
	l.head.next.prev = l.head
	return l.head
}

func (l *DoublyLinkedList) Remove(item *ListNode) {
	if l.head != nil {
		current := &(l.head)
		for *current != nil {
			if (*current).next == item {

				(*current).next = (*current).next.next
				if (*current).next == nil {
					l.tail = *current
				} else {
					(*current).next.prev = (*current)
				}
				break
			}
			current = &((*current).next)
		}
	}
}
