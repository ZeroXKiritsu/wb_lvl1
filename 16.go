package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name     string
	LastName string
	Age      int
}

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	slice := []int{-25, -27, 13, 19, 15, 24, -21, 32}
	
	person := []Person{
		{"Karen", "Gasparyan", 24},
		{"Alex", "Govard", 36},
		{"Unknown", "Person", 15655},
	}
	sort.Sort(ByName(person))
	fmt.Println(person)

	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)
}
