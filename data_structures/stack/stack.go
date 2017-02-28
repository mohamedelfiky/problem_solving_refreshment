package stack

type StackItem struct {
	value interface{}
	next  *StackItem
}

type Stack struct {
	top  *StackItem
	size int
}

func (s *Stack) Push(value interface{}) *StackItem {
	s.top = &StackItem{value: value, next: s.top}
	s.size++
	return s.top
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	top := s.top
	s.top = s.top.next
	s.size--
	return top.value
}
