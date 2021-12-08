package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stopGoroutine()
}

func stopGoroutine() {
	var num time.Duration
	fmt.Println("Enter the number of seconds: ")
	fmt.Scan(&num)

	chann := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	writer := func(ctx context.Context) <-chan int {
		i := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(chann)
					return
				case chann <- i:
					i++
					fmt.Println("working")
				}

				time.Sleep(500 * time.Millisecond)
			}
		}()
		return chann
	}

	go func() {
		time.Sleep(num * time.Second)
		cancel()
	}()

	for i := range writer(ctx) {
		fmt.Println(i)
	}
}
