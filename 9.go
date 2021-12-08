package main

import "fmt"

func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	arr := []int{2, 4, 6, 8, 10}
	go readChannel(c0, arr)
	go squaredChannel(c0, c1)
	for i := range c1 {
		fmt.Println(i)
	}
}

func squaredChannel(ch chan int, result chan int) {
	for i := range ch {
		result <- 2 * i
	}
	close(result)
}

func readChannel(ch chan int, arr []int) {
	for _, v := range arr {
		ch <- v
	}
	close(ch)
}
