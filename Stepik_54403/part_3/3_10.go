package main

import (
	"fmt"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	var summ int
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
			case x := <-arguments:
				summ += x
			case <-done:
				output <- summ
				return
			}
		}
	}()

	return output
}

func main() {

	fmt.Println("1")
}

// Вам необходимо написать функцию calculator следующего вида:

// func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
// В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

// Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о
// завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

// Функция calculator должна быть неблокирующей, сразу возвращая управление.

// Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.
