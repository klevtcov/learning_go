package main

import (
	"fmt"
)

// Structure - data encapsulation
// Структуры представляют тип данных, определяемый разработчиком и служащий для представления каких-либо объектов. Структуры содержат набор полей, которые представляют различные атрибуты объекта. Для определения структуры применяются ключевые слова type и struct:
// https://metanit.com/go/tutorial/4.2.php

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	// c := Car{"chevy", 1, 2} - {chevy 1 2}
	c := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2,
	}

	c.Print()
	c.Drive()
	fmt.Println(c.GetName())

	// var c1 Car

	// fmt.Println(c)

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
