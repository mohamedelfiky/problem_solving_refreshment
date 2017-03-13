package priority_queue

type Item struct {
	Order     int
	Operation func()
}

type PriorityQueue struct {
	queue []Item
}

func (p *PriorityQueue) Push(order int, operation func()) {
	p.queue = append(p.queue, Item{order, operation})
	p.heapifyUp(p.Size() - 1)
}

func (p *PriorityQueue) Pop() Item {
	if p.Size() == 0 {
		panic("Queue has no items")
	}
	first := p.queue[0]
	p.queue[0] = p.queue[len(p.queue)-1]
	p.queue = p.queue[:len(p.queue)-1]
	p.heapifyDown(0)
	return first
}

func (p *PriorityQueue) Size() int {
	return len(p.queue)
}

func (p *PriorityQueue) heapifyUp(index int) {
	for parent(index) != index && p.queue[index].Order < p.queue[parent(index)].Order {
		swap(&p.queue, index, parent(index))
		index = parent(index)
	}
}

func (p *PriorityQueue) heapifyDown(index int) {
	if left(index) < p.Size() && p.queue[index].Order > p.queue[left(index)].Order {
		swap(&p.queue, index, left(index))
		p.heapifyDown(left(index))
	}
	if right(index) < p.Size() && p.queue[index].Order > p.queue[right(index)].Order {
		swap(&p.queue, index, right(index))
		p.heapifyDown(right(index))
	}
}

func parent(index int) int { return (index - 1) / 2 }
func left(index int) int   { return (2 * index) + 1 }
func right(index int) int  { return (2 * index) + 2 }

func swap(arr *[]Item, i1, i2 int) {
	tmp := (*arr)[i1]
	(*arr)[i1] = (*arr)[i2]
	(*arr)[i2] = tmp
}
