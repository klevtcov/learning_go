package main

import "fmt"

func main() {

	keys := []string{}

	// таблица, ключ - строка, значение - массив чисел
	carta := map[string][]int{
		"1": {1, 2},
		"2": {2, 10},
	}

	for key := range carta {
		keys = append(keys, key)
		fmt.Println(key) // 1 2 - порядок может быть любой
	}

	for key, value := range carta {
		fmt.Println(key, value) // 1 [1 2], 2 [2 10]
	}

	fmt.Println(keys) // [1 2]

	fmt.Println(carta["2"]) // [2 10]

	h := make(map[int]string, 50) // можно сразу задать размер карты, это ускорит работу
	fmt.Println(h)                // map[]

}
