package sorting

import "time"

func SleepSort(arr []int) []int {
	ch := make(chan int)
	for _, num := range arr {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(num)
	}

	result := []int{}
	for i := 0; i < len(arr); i++ {
		result = append(result, <-ch)
	}
	return result
}
