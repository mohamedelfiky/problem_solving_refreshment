package circular_queue

const MAX_SIZE int = 3 // small for testing

type CircularQueue struct {
	queue             [MAX_SIZE]interface{}
	front, rear, size int
}

func NewCircularQueue() *CircularQueue {
	return &CircularQueue{front: -1, rear: -1}
}

func (l *CircularQueue) Push(value interface{}) {
	if !l.IsFull() {
		l.rear = (l.rear + 1) % MAX_SIZE
		l.queue[l.rear] = value
		if l.front == -1 {
			l.front = 0
		}
		l.size++
	}
}

func (l *CircularQueue) Pop() interface{} {
	if !l.IsEmpty() {
		front := l.front
		l.front = (front + 1) % MAX_SIZE
		l.size--
		return l.queue[front]
	}
	return nil
}

func (l *CircularQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *CircularQueue) IsFull() bool {
	return l.size == MAX_SIZE
}
