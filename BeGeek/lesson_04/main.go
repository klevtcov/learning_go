package main

import (
	"fmt"
	"sort"
)

func main() {

	// Массивы
	// var array [3]string // массив из 3-х элементов, каждый из которых строка
	// fmt.Println(array) // [  ] три пустых значения
	// array[1] = "be"
	// fmt.Println(array)      // [ be ] - пусто, строка, пусто
	// fmt.Println(len(array)) // 3

	// array := [3]string{} // при явном указании длинны массива, в дальнейшем изменить будет нельзя - память выделяется под три элемента, можно только менять их значения
	// array := [3]bool{} /// False False False
	// array := [3]int{1, 3, 4}

	// fmt.Println(len(array)) // 3
	// fmt.Println(array)      // [1 3 4]

	// for _, i := range array {
	// 	fmt.Println(i) // 1 3 4
	// }

	// Срезы
	// array := []int{1, 2, 3}
	// fmt.Println(len(array)) // 3

	// array = append(array, 321)
	// fmt.Println(len(array)) // 4
	// fmt.Println(array)      // [1 2 3 321]

	array := make([]int, 10)
	fmt.Println(len(array)) // 10
	fmt.Println(array)      // [0 0 0 0 0 0 0 0 0 0]

	array[2] = 12
	fmt.Println(array) // [0 0 12 0 0 0 0 0 0 0]
	sort.Ints(array)
	fmt.Println(array) // [0 0 0 0 0 0 0 0 0 12]

}
