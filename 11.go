package main

import "fmt"

func main() {
	slice1 := []string{"foo", "bar", "test"}
	slice2 := []string{"foo", "bar"}
	fmt.Printf("%+v\n", intersection(slice1, slice2))
}

func intersection(a []string, b []string) (result []string) {
	// взаимодействие с наименьшим списком в первую очередь может быть потенциально быстрее...но не намного, худший случай тот же самый
	small, big := a, b
	if len(a) > len(b) {
		small = b
		big = a
	}

	ok := false
	for i, v := range small {
		for index, val := range big {
			// получить будущие значения индекса
			future := i + 1
			futureIndex := index + 1
			if v == val {
				result = append(result, val)
				if future < len(small) && futureIndex < len(big) {
					// если будущие значения не совпадают, то это конец пересечения
					if small[future] != big[futureIndex] {
						ok = true
					}
				}
				//чтобы каждый раз не взаимодействовать со всем списком, поэтому удаляем части, которые мы уже зациклили, таким образом ускорится каждый проход
				big = big[:index+copy(big[index:], big[index+1:])]
				break
			}
		}
		if ok {
			break
		}
	}
	return result
}
