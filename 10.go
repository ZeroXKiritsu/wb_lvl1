package main

import "fmt"

func main() {
	result := AppendByGroup(-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)
	fmt.Println(result)
}

func AppendByGroup(num ...float32) map[int][]float32 {
	result := make(map[int][]float32)

	for _, v := range num {
		category := int(v) - (int(v) % 10)
		if _, ok := result[category]; ok {
			result[category] = append(result[category], v)
		} else {
			result[category] = []float32{v}
		}
	}
	return result
}