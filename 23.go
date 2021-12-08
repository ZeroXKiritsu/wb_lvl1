package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(arr)
	fmt.Println(remove(arr, 2))
}

func remove(slice []int, s int) []int {
	// Копируем последний элемент в индекс i
	slice[s] = slice[len(slice)-1]
	// В последний элемент записываем ноль
	slice[len(slice)-1] = 0
	// срезаем массив
	slice = slice[:len(slice)-1]
	return slice
} 