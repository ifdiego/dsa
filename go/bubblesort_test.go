package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	BubbleSort(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("bubblesort failed! expected %v, but got %v", expected, input)
	}
}
