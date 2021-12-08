package main

import "fmt"

func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, v := range s {
		m[v] = true
	}
	return len(s) == len(m)
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(isUnique(str1))
	fmt.Println(isUnique(str2))
	fmt.Println(isUnique(str3))
}
