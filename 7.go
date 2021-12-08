package main

import (
	"fmt"
	"sync"
)

func main() {
	Map := make(map[int]string)

	mu := &sync.Mutex{}

	wg := &sync.WaitGroup{}

	keyArr := []int{2,4,6,8,10}

	valArr := []string{"two", "four", "six", "eight", "ten"}

	for i := 0; i < len(keyArr); i++ {
		wg.Add(1)
		go func(m map[int]string, keys []int, val []string, index int, mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			m[keys[index]] = val[index]
			mu.Unlock()
		}(Map, keyArr, valArr, i, mu, wg)
	}
	wg.Wait()
	fmt.Println(Map)	
}