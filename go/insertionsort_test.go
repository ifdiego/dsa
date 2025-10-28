package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	nums := InsertionSort(input)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("insertion sort failed! expected %v, but got %v", expected, nums)
	}
}
