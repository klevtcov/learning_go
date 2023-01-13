package main

import (
	"fmt"
)

// Structure - data encapsulation
// Структуры представляют тип данных, определяемый разработчиком и служащий для представления каких-либо объектов. Структуры содержат набор полей, которые представляют различные атрибуты объекта. Для определения структуры применяются ключевые слова type и struct:
// https://metanit.com/go/tutorial/4.2.php

type Car struct {
}

func main() {
	c := Car{}
	var c1 Car

	fmt.Println()

}

// arrays

// func todo() {
// 	// var arr []int
// 	arr := []int{1, 2, 3, 4}
// 	arr2 := []string{"hi", "my", "name"}
// 	// fmt.Println(arr, arr2) - [1 2 3 4] [hi my name]
// 	// fmt.Println(arr) - [1 2 3 4]
// 	// arr2 = append(arr2, "is Sergey")
// 	// fmt.Println(arr, arr2) - [1 2 3 4] [hi my name is Sergey]
// 	fmt.Println(arr, arr2)
// }
