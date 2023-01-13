package main

import (
	"fmt"
)

func main() {

	fmt.Println()

}

// // Int

// var m1 int
// m1 = 2
// fmt.Println(m1)

// var (
// 	m1 = 2
// 	m2 = 3
// )
// fmt.Println(m1 + m2)

// var m1 int32
// var m2 int64
// fmt.Println(m1 + m2) - Error, нелльзя складывать разные типы, нужно приведение к типу

// var m1 int32
// var m2 int64
// fmt.Println(int64(m1) + m2) - ответ 0, т.к. по умолчанию переменная создаётся с 0 значением

// m1 := 2 - объявление и присвоение в одном операторе. го сам распознаёт тип
// m2 := 3
// fmt.Println(m1 + m2)

// String

// var m1 string
// m1 = "my name"

// m1 := "my name"
// m2 := "name"
// fmt.Println(strings.Contains(m1, m2)) проверяем вхождение строки

// m1 := "my name"
// m2 := "name"
// fmt.Println(strings.ReplaceAll(m1, "m", "NO")) - NOy naNOe. замена символов в строке

// m1 := "my name"
// m2 := "name"
// fmt.Println(strings.Split(m1, " ")) - [my name] , можно обращаться через [0]

// m1 := "my name"
// m2 := "name"
// fmt.Println(m1 + m2) - my namename конкатенация строк

// Pointers

// func swap(m1, m2 *int) {
// 	var temp int
// 	temp = *m2
// 	*m2 = *m1
// 	*m1 = temp
// }

// Pointers
// m1 := 3
// ptr := &m1

// // fmt.Println(ptr) - 0xc0000a6058 адрес в памяти
// fmt.Println(*ptr) - 3 (*)если хотим обратиться к объекту по адресу

// m1 := 3
// m2 := 4
// swap(&m1, &m2)
// fmt.Println("m1 - ", m1, " m2 - ", m2)
