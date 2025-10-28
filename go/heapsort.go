package sorting

import "math"

func HeapSort(arr []int) {
	size := len(arr)
	for i := math.Floor(float64(size/2)) - 1; i >= 0; i-- {
		Heapify(arr, size, int(i))
	}

	j := size - 1
	for j >= 1 {
		arr[0], arr[j] = arr[j], arr[0]
		Heapify(arr, j, 0)
		j--
	}
}

func Heapify(arr []int, size, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, size, largest)
	}
}
