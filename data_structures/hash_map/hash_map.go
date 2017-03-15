package hash_map

import "github.com/mohamedelfiky/problem_solving_refreshment/data_structures/linked_list"

type MapItem struct {
	Key   string
	Value interface{}
}

// linked list implement of HashMap
type HashMap struct {
	harr []*linked_list.LinkedList
}

const TABLE_SIZE int = 150

func New() HashMap {
	return HashMap{harr: make([]*linked_list.LinkedList, TABLE_SIZE)}
}

func (h *HashMap) Put(key string, value interface{}) {
	index := hindex(key)
	obj := MapItem{Key: key, Value: value}
	if h.harr[index] == nil {
		l := new(linked_list.LinkedList)
		l.Push(obj)
		h.harr[index] = l
	} else {
		h.harr[index].Push(obj)
	}
}

func (h *HashMap) Get(key string) interface{} {
	index := hindex(key)
	if h.harr[index] != nil {
		result := h.harr[index].SearchBy(func(i interface{}) bool { return i.(*linked_list.ListNode).Value.(MapItem).Key == key })
		if result != nil {
			return result.Value
		}
	}
	return nil
}

func (h *HashMap) Remove(key string) {
	index := hindex(key)
	if h.harr[index] != nil {
		if h.harr[index].Head != nil && h.harr[index].Head.Value.(MapItem).Key == key {
			h.harr[index].Remove(h.harr[index].Head)
			return
		}
		result := h.harr[index].SearchBy(func(i interface{}) bool {
			next := i.(*linked_list.ListNode).Next()
			return next != nil && next.Value.(MapItem).Key == key
		})
		if result != nil {
			result.RemoveNext()
		}
	}
}

func hindex(key string) int {
	sum, arr := 0, []rune(key)
	for i := 0; i < len(arr); i++ {
		sum += int(arr[i])
	}
	return sum % TABLE_SIZE
}
