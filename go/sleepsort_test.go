package sorting

import (
	"reflect"
	"testing"
)

func TestSleepSort(t *testing.T) {
	input := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	nums := SleepSort(input)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("sleep sort failed! expected %v, but got %v", expected, nums)
	}
}
