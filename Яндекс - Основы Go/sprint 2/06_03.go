package ma063

import "fmt"

var Global = 5

func changing(g int) {
	g = 7
	fmt.Println(g)
	g = 5
}

func main() {
	defer changing(Global)
}

// Директиву defer обычно применяют для освобождения ресурсов — например, чтобы закрыть каналы, файлы,
// сетевые соединения после окончания работы. Об использовании системных ресурсов расскажем в следующих
// темах курса, а пока рассмотрим в качестве ресурса глобальную переменную, доступную всем функциям программы:
// var Global = 5
// Используя defer, напишите функцию, которая меняет эту переменную и выводит на экран её новое значение.
// Потом она должна вернуть всё как было.

// package main

// import (
//     "fmt"
// )
// // глобальная переменная равна 5
// var Global = 5

// func UseGlobal() {
//     defer func(checkout int) {
//         Global = checkout // присваиваем Global значение аргумента
//     }(Global) // копируем значение Global в аргументы функции
//     Global = 42 // Изменяем Global
//     fmt.Println(Global)
//     // Здесь будет вызвана отложенная функция
// }

// func main() {
//     fmt.Println(Global)
//     UseGlobal()
//     fmt.Println(Global)
// }
