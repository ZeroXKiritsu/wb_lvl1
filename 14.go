package main

import "fmt"

func main() {
	type1 := "hello world"
	type2 := 666
	type3 := 24.5
	type4 := make(chan int, 100)
	type5 := false
	TypeOf(type1)
	TypeOf(type2)
	TypeOf(type3)
	TypeOf(type4)
	TypeOf(type5)
}

func TypeOf(v interface{}) {
	fmt.Printf("%T\n", v)
}
