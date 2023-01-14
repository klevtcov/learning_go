package main

import (
	"fmt"
	"time"
)

func main() {

	// var c chan int // <nil> каанл объявлен, но не создан
	// <name> chan <datatype>

	c := make(chan int)
	// 0xc0000200c0 - адрес созданого канала

	// send something to chanel. должен соблюдаться тип данных
	go func() {
		c <- 1
	}()
	// take from chanel
	val := <-c

	fmt.Println(val)

	go func() {
		c <- 2
	}()
	time.Sleep((time.Second * 2))
	val = <-c
	fmt.Println(val)

	fmt.Println(c)
}
