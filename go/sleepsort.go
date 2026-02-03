package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepsort(arr []int) {
	var wg sync.WaitGroup
	for _, n := range arr {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(i)
		}(n)
	}
	wg.Wait()
}

func main() {
	numbers := []int{7, 42, 19, 88, 3, 56, 91, 24, 65, 11}
	sleepsort(numbers)
}
