package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	input := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	HeapSort(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("heap sort failed! expected %v, but got %v", expected, input)
	}
}
