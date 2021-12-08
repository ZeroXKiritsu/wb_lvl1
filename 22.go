package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := int64(math.Pow(2, 24))
	b := int64(math.Pow(2, 22))
	var choice string
	fmt.Scanf("%v", &choice)
	switch choice {
	case "+":
		fmt.Println("result:", add(a, b))
	case "-":
		fmt.Println("result:", sub(a, b))
	case "*":
		fmt.Println("result:", mul(a, b))
	case "/":
		fmt.Println("result:", div(a, b))
	}
}

func mul(a, b int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
}

func div(a, b int64) *big.Int {
	return new(big.Int).Div(big.NewInt(a), big.NewInt(b))
}

func sub(a, b int64) *big.Int {
	return new(big.Int).Sub(big.NewInt(a), big.NewInt(b))
}

func add(a, b int64) *big.Int {
	return new(big.Int).Add(big.NewInt(a), big.NewInt(b))
}
