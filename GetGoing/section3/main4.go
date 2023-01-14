package main

import "fmt"

type Car struct {
	Model string
}

func main() {

	c := make(chan *Car, 3)

	go func() {
		c <- &Car{"1 car"}
		c <- &Car{"2 car"}
		c <- &Car{"3 car"}
		c <- &Car{"4 car"}
		close(c)
	}()

	for i := range c {
		fmt.Println(i.Model)
	}

	// c := make(chan int, 3)
	// go func() {
	// 	c <- 1
	// 	c <- 2
	// 	c <- 3
	// 	c <- 4
	// 	close(c)
	// }()

	// for i := range c {
	// 	fmt.Println(i)
	// }

	fmt.Println("Fin")
}
