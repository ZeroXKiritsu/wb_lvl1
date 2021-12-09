package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopGoroutineByDeadline()
}

func stopGoroutineByDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("overslept")
			}
		}
	}()
}

func stopGoroutineByTime() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}()
}

func stopGoroutineByWaitGroup() {
	chann := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		for {
			select {
			case <-chann:
				return
			}
		}
	}()
}

func stopGoroutineByClose() {
	chann := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		close(chann)
	}()
}

func stopGoroutineByContext() {
	chann := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	writer := func(ctx context.Context) <-chan int {
		i := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case chann <- i:
					i++
					fmt.Println("working")
				}
			}
		}()
		return chann
	}

	for i := range writer(ctx) {
		if i == 3 {
			break
		}
	}
}
