package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	mu    *sync.Mutex
}

func NewCounter() Counter {
	counter := Counter{}
	counter.count = 0
	counter.mu = &sync.Mutex{}
	return counter
}

func (c *Counter) Add(i int) {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	fmt.Println(i, c.count)
}

func main() {
	counter := NewCounter()
	go counter.Add(0)
	go counter.Add(1)
	go counter.Add(2)
	go counter.Add(3)
	go counter.Add(4)
	go counter.Add(5)
	time.Sleep(time.Second)
}
