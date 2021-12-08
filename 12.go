package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	Map := make(map[string]struct{})
	var set []string

	for i := range arr {
		Map[arr[i]] = struct{}{}
	}

	for i := range Map {
		set = append(set, i)
	}
	fmt.Println(arr)
	fmt.Println(set)
}