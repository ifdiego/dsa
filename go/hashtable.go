package hashtable

import (
	linkedlist "github.com/ifdiego/algorithms/linked-list"
)

type HashTable struct {
	size int
	data []*linkedlist.LinkedList[string]
}

func CreateHashTable(size int) *HashTable {
	return &HashTable{
		size: size,
		data: make([]*linkedlist.LinkedList[string], size),
	}
}

func (h *HashTable) hash(value string) int {
	sum := 0
	for _, c := range value {
		sum += int(c)
	}
	return sum % h.size
}

func (h *HashTable) Insert(value string) {
	index := h.hash(value)
	if h.data[index] == nil {
		l := &linkedlist.LinkedList[string]{}
		l.Equals = func(a, b string) bool { return a == b }
		h.data[index] = l
	}
	h.data[index].Append(value)
}

func (h *HashTable) Search(value string) *string {
	index := h.hash(value)
	if h.data[index] != nil {
		node := h.data[index].Search(value)
		if node != nil {
			return &node.Value
		}
	}
	return nil
}
