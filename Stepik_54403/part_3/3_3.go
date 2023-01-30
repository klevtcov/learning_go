package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fn := func(i uint) uint {
		str := strconv.Itoa(int(i))
		runes := []rune(str)
		fmt.Println(runes)
		var result []rune
		for _, val := range runes {
			val -= '0'
			if val%2 == 0 && val != '0' {
				result = append(result, val+'0')
			}
		}
		fmt.Println(result)
		s := string(([]rune(result)))
		si, _ := strconv.Atoi(s)
		// res, _ := strconv.Atoi(result)
		if si == 0 {
			return uint(100)
		}
		return uint(si)
	}
	fmt.Println(fn(272178))

	// src = strings.Map(func(r rune) rune {return unicode.ToUpper(r)}, src)

	// вычитаем из ASCII-символа цифры значение ASCII-символа нуля, получаем собственно цифру.
	// var s string
	// fmt.Scan(&s)
	// for _, c := range s {
	// 	c -= '0'
	// 	fmt.Print(c * c)
	// }

}

// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main
// в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные
// цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
//
//

//
// пример работы функции - можем вызвать функцию и передавть в неё две другие и два аргумента. логика прописана внутри функции
// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	fmt.Println(sum(up10, del2, a, b))
// }
// func sum(fa func(int) int, fb func(int) int, a, b int) int {
// 	return fa(a) + fb(b)
// }
// func up10(x int) int {
// 	return x * 10
// }
// func del2(c int) int {
// 	return c / 2
// // }
//
//
// x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
// fmt.Printf("%T", x)
