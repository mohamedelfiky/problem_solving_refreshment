package linked_list

type ListNode struct {
	Value interface{}
	next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) Push(value interface{}) *ListNode {
	node := &ListNode{Value: value}
	if l.IsEmpty() {
		l.Head = node
	} else {
		last := &(l.Head)
		for (*last).next != nil {
			last = &((*last).next)
		}
		(*last).next = node
	}
	return node
}

func (l *LinkedList) Insert(afterMe *ListNode, value interface{}) *ListNode {
	node := &ListNode{Value: value}
	node.next = afterMe.next
	afterMe.next = node
	return node
}

func (l *LinkedList) Search(value interface{}) *ListNode {
	if !l.IsEmpty() {
		return l.SearchBy(func(item interface{}) bool {
			return item.(*ListNode).Value == value
		})
	}
	return nil
}

func (l *LinkedList) SearchBy(comparer func(item interface{}) bool) *ListNode {
	if !l.IsEmpty() {
		current := l.Head
		for current != nil {
			if comparer(current) {
				return current
			}
			current = current.next
		}
	}
	return nil
}

func (l *LinkedList) PushFirst(value interface{}) {
	l.Head = &ListNode{Value: value, next: l.Head}
}

func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

func (n *ListNode) Next() *ListNode {
	return n.next
}

func (l *LinkedList) Remove(item *ListNode) {
	if !l.IsEmpty() {
		if l.Head == item {
			l.Head = l.Head.next
			return
		}
		current := l.Head
		for current != nil {
			if current.next == item {
				current.RemoveNext()
				break
			}
			current = current.next
		}
	}
}

func (n *ListNode) RemoveNext() {
	if n != nil && n.next != nil {
		n.next = n.next.next
	}
}
