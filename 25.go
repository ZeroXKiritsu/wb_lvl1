package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	sleep(3)
	fmt.Println("end")
}

func sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}
