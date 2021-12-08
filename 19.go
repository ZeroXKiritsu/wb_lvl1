package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(str)
	str = Reverse(str)
	fmt.Println(str)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
