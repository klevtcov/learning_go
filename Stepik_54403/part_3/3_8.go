package m3_8

import "fmt"

func task(c chan int, N int) {
	c <- N + 1
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var prev string
	for s := range inputStream {
		if s != prev {
			outputStream <- s
		}
		prev = s
	}
}

func main() {

	// done := make(chan struct{})
	// go myFunc(done)
	// <-done
	fmt.Println("2")
}

// Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.

// Функция work() ничего не принимает и не возвращает.
//
// done := make(chan struct{})
// go func(d chan struct{}) {
// 	work()
// 	close(d)
// }(done)
// <-done
//
// создаём канал пустых структур;
// создаём анонимную функцию и сразу же запускаем её в виде горутины;
// внутри горутины вызываем work() и закрываем канал;
// основной поток всё это время после запуска горутины ждёт закрытия канала. work() может
// работать сколь угодно долго, и основной поток в любом случае дождётся её завершения.
//
//
// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().

// Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

//
// --------------

// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
// Функция должна называться task2().
//
// -------------------
//
//
// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап
// конвейера только если оно отличается от того, что пришло ранее.
// Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
// во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения,
// которые не повторяются подряд. Не забудьте закрыть канал ;)
// Функция должна называться removeDuplicates()
