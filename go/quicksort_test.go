package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	QuickSort(input, 0, len(input)-1)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("quick sort failed! expected %v, but got %v", expected, input)
	}
}
