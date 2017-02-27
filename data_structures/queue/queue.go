package queue

type QueueItem struct {
	value interface{}
	next  *QueueItem
}

type Queue struct {
	size       int
	head, rear *QueueItem
}

func (l *Queue) Push(value interface{}) *QueueItem {
	item := &QueueItem{value: value}
	if l.IsEmpty() {
		l.head = item
		l.size = 1
	} else {
		l.rear.next = item
		l.size++
	}
	l.rear = item
	return item
}

func (l *Queue) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	} else {
		head := l.head
		l.head = l.head.next
		l.size--
		if l.size == 0 {
			l.rear = nil
		}
		return head.value
	}
}

func (l *Queue) IsEmpty() bool {
	return l.head == nil
}
