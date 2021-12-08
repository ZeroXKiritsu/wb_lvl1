package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(in chan interface{}, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for i := range in {
		fmt.Printf("worker: %d, value: %v", num + 1, i)
	}
	
}

func main() {
	var num int
	wg := &sync.WaitGroup{}
	chann := make(chan interface{})
	end := make(chan int)
	fmt.Print("enter count of workers")
	fmt.Scanf("%d\n", &num)
	wg.Add(num)
	go func(ch chan int) {
		time.Sleep(time.Second * 1)
		ch <- 0
	}(end)
	for i := 0; i < num; i++ {
		go workers(chann, wg, i)
	}
	for {
		select {
		case chann <- "1":
		case chann <- "2":
		case chann <- "3":
		case <- end:
			return
		}
	}
	close(chann)
	close(end)
}