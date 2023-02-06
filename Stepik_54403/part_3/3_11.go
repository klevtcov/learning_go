package main

import (
	"fmt"
	"sync"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	chan1 := make([]int, n)
	chan2 := make([]int, n)

	wg := new(sync.WaitGroup)

	go func() {
		defer close(out)

		for i := 0; i < n; i++ {
			x := <-in1
			y := <-in2
			chan1[i] = x
			chan2[i] = y
		}

		wg.Add(n * 2)

		for i := 0; i < n; i++ {

			go func(i int) {
				chan1[i] = fn(chan1[i])
				wg.Done()
			}(i)

			go func(i int) {
				chan2[i] = fn(chan2[i])
				wg.Done()
			}(i)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			out <- chan1[i] + chan2[i]
		}

	}()

}

func main() {

	fmt.Println("2")
}

// Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

// Описание ее работы:

// n раз сделать следующее

// прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
// вычислить f(x1) + f(x2)
// записать полученное значение в out
// Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

// Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

// Формат ввода:

// количество итераций передается через аргумент n.
// целые числа подаются через аргументы-каналы in1 и in2.
// функция для обработки чисел перед сложением передается через аргумент fn.
// Формат вывода:

// канал для вывода результатов передается через аргумент out.
