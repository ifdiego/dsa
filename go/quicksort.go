package sorting

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pointer := Rearrange(arr, start, end)
	QuickSort(arr, start, pointer-1)
	QuickSort(arr, pointer+1, end)
}

func Rearrange(arr []int, start, end int) int {
	pivot := arr[end]
	pointer := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[pointer], arr[i] = arr[i], arr[pointer]
			pointer++
		}
	}
	arr[pointer], arr[end] = arr[end], arr[pointer]
	return pointer
}
