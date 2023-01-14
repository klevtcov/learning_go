package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func super_heavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}

func main() {

	go heavy()
	go super_heavy()
	fmt.Println("Fin")
	select {}

	// time.Sleep(time.Second * 5)
	// select {} - ждет выполнения всех горутин
}
