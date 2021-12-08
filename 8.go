package main

import "fmt"

func main() {
	fmt.Println("set to 1: %b and %d\n", setBitOne(1, 2))
	fmt.Println("set to 0: %b and %d\n", setBitZero(1, 2))
}

func setBitOne(num int64, i int64) int64 {
	return (num | (1 << (i - 1)))
}

func setBitZero(num int64, i int64) int64 {
	return (num & ^(1 << (i - 1)))
}