package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {

	// swith case
	// select for channel async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("recived")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(1)
		}
	}
}

func main() {

	c := make(chan int, 2)
	quits := make(chan struct{})

	go func() {
		c <- 1
		quits <- struct{}{}
	}()

	go Select(c, quits)

	select {} // пустому селекту не передается имя канала, который он должен слыушть, поэтому он зависает в ожидании и не может дождаться сигнала. поэтому тормозится выполнение потока
}
