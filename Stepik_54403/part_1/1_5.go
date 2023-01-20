package m1_5

import "fmt"

func main() {

	fmt.Println("I like Go!")
}

// fmt.Println("I like Go!")
// fmt.Println("I like Go!")

// //Так как строки в Go хранятся в виде байтов нужно понимать что такой код не выведет символ:
// fmt.Println("Hello Go"[0])         // вывод: 72
// fmt.Println(string("Hello Go"[0])) // вывод: H

// Чтение с консоли нескольких переменных
// fmt.Scan(&a, &b, &c)

// Чтение с консоли
// var name string
// var age int
// fmt.Print("Введите имя: ")
// fmt.Scan(&name)
// fmt.Print("Введите возраст: ")
// fmt.Scan(&age)

// fmt.Println(name, age)

// fmt.Println("")

// Напишите программу, которая последовательно делает следующие операции с введённым числом:

// Число умножается на 2;
// Затем к числу прибавляется 100.
// var a int
// fmt.Scan(&a)
// // здесь ваш код
// fmt.Println((a * 2) + 100)

// var a, b int
// fmt.Scan(&a) // считаем переменную 'a' с консоли
// fmt.Scan(&b) // считаем переменную 'b' с консоли

// a = a * a
// b = b * b
// c := a + b
// fmt.Println(c) // Петя торопился в школу и неправильно написал программу, которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.

// По данному целому числу, найдите его квадрат.
// Формат входных данных
// На вход дается одно целое число.
// Формат выходных данных
// Программа должна вывести квадрат данного числа.
// var a int
// fmt.Scan(&a)
// fmt.Println(a * a)

// 	Дано натуральное число, выведите его последнюю цифру.
// 	Формат входных данных
// 	На вход дается натуральное число N, не превосходящее 10000.
// 	Формат выходных данных
// 	Выведите одно целое число - ответ на задачу.
// var n int
// fmt.Scan(&n)
// fmt.Println(n % 10)

// Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.
// Формат выходных данных
// Выведите одно целое число - число десятков.
// var n int
// fmt.Scan(&n)
// fmt.Println((n % 100) / 10)

// Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
// Входные данные
// На вход программе подается целое число d (0 < d < 360).
// Выходные данные
// Выведите на экран фразу:
// It is ... hours ... minutes.
// Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
// package main

// import "fmt"

// func main() {
// 	var d, h, m int
// 	fmt.Scan(&d)
// 	h = (d * 2) / 60
// 	m = (d * 2) % 60
// 	fmt.Println("It is", h, "hours", m, "minutes.")
// }