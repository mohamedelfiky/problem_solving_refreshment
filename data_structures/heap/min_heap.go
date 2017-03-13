package heap

type MinHeap struct {
	harr []int
}

func (heap *MinHeap) Push(value int) {
	heap.harr = append(heap.harr, value)
	heap.heapifyUp(len(heap.harr) - 1)
}

func (heap *MinHeap) Pop() int {
	if len(heap.harr) == 0 {
		panic("Heap has no values")
	}
	first := heap.harr[0]
	heap.harr[0] = heap.harr[len(heap.harr)-1]
	heap.harr = heap.harr[:len(heap.harr)-1]
	heap.heapifyDown(0)
	return first
}

func (heap *MinHeap) Size() int {
	return len(heap.harr)
}

func (heap *MinHeap) heapifyUp(index int) {
	for parent(index) != index && heap.harr[index] < heap.harr[parent(index)] {
		swap(&heap.harr, index, parent(index))
		index = parent(index)
	}
}

func (heap *MinHeap) heapifyDown(index int) {
	if left(index) < heap.Size() && heap.harr[index] > heap.harr[left(index)] {
		swap(&heap.harr, index, left(index))
		heap.heapifyDown(left(index))
	}
	if right(index) < heap.Size() && heap.harr[index] > heap.harr[right(index)] {
		swap(&heap.harr, index, right(index))
		heap.heapifyDown(right(index))
	}
}

func parent(index int) int { return (index - 1) / 2 }
func left(index int) int   { return (2 * index) + 1 }
func right(index int) int  { return (2 * index) + 2 }

func swap(arr *[]int, i1, i2 int) {
	tmp := (*arr)[i1]
	(*arr)[i1] = (*arr)[i2]
	(*arr)[i2] = tmp
}
