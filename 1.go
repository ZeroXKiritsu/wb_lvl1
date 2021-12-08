package main

import "fmt"

type Human struct {
	name     string
	lastName string
	age      uint
}

type Action struct {
	*Human
}

func main() {
	action := &Action{Human: &Human{"Karen", "Gasparyan", 24}}
	fmt.Printf("%s %s %v", action.name, action.lastName, action.age)
}