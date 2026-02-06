package main

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := arr[0 : len(arr)/2]
	right := arr[len(arr)/2 : len(arr)]
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	sorted := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			value := left[0]
			left = left[1:]
			sorted = append(sorted, value)
		} else {
			value := right[0]
			right = right[1:]
			sorted = append(sorted, value)
		}
	}
	return append(append(sorted, left...), right...)
}
