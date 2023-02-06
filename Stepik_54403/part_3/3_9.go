package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(<-chan int) 

	go func() {
		defer close(output)

		select {
		case <-firstChan:
			a := <-firstChan
		case <-tick2:
			fmt.Println("Получено значение из второго канала")
		}
	}
}



tick1 := time.After(time.Second)
tick2 := time.After(time.Second * 2)
select {
case <-tick1:
	fmt.Println("Получено значение из первого канала")
case <-tick2:
	fmt.Println("Получено значение из второго канала")
// Блок default выполнится раньше блока case - 1 секунда слишком много для Go
default:
	fmt.Println("Действие по умолчанию")
}

func main() {

	fmt.Println("2")
}

// Вам необходимо написать функцию calculator следующего вида:

// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
// Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

// в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
// в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
// в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
// Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

// После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.
