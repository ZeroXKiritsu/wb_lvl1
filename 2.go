package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{2,4,6,8,10}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(num int, wg *sync.WaitGroup)  {
			defer wg.Done()
			fmt.Printf("%d\n", num * num)
		}(arr[i], wg)
	}
	wg.Wait()
}