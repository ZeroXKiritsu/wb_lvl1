package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	var duration time.Duration
	fmt.Println("Enter the number of seconds: ")
	fmt.Scan(&duration)
	wg.Add(1)
	go RwWaiter(duration, wg)
	wg.Wait()
}

func RwWaiter(duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(duration * time.Second))
	chann := make(chan interface{})

	go func(ch chan interface{}, ctx context.Context) {
		data := []interface{}{"Yes", "No", "Maybe"}
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <- ticker.C:
				ch <- data[rand.Intn(len(data))]
			case <- ctx.Done():
				ticker.Stop()
				return
			}
		} 
	}(chann, ctx)

	go func (ch chan interface{}, ctx context.Context)  {
		for {
			select {
			case val := <- ch:
				fmt.Printf("Readed value %v from channel\n", val)
			case <-ctx.Done():
				return				
			}
		}
	}(chann, ctx)
	<-ctx.Done()
}