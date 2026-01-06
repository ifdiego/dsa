package main

type HashTable struct {
	data []map[string]string
	size int
}

func NewHashTable(size int) *HashTable {
	data := make([]map[string]string, size)
	for i := range data {
		data[i] = make(map[string]string)
	}
	return &HashTable{
		data: data,
		size: size,
	}
}

func (h *HashTable) hash(key string) int {
	value := 0
	for _, c := range key {
		value += int(c)
	}
	return value % h.size
}

func (h *HashTable) Insert(key, value string) {
	index := h.hash(key)
	h.data[index][key] = value
}

func (h *HashTable) Search(key string) (string, bool) {
	index := h.hash(key)
	value, exists := h.data[index][key]
	return value, exists
}

func (h *HashTable) Remove(key string) {
	index := h.hash(key)
	delete(h.data[index], key)
}
