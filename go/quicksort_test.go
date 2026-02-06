package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{12, 6, 18, 19, 21, 11, 3, 5, 4, 24, 8}
	expected := []int{3, 4, 5, 6, 8, 11, 12, 18, 19, 21, 24}
	result := QuickSort(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Quicksort failed! Expected %v, but got %v", expected, result)
	}
}
