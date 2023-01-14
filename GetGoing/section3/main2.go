package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait group
	wg := &sync.WaitGroup{}
	// add. done and wait functionality

	wg.Add(2) // задаём количество "ок" от горутин, после которых сработает код ниже отметки wg.Wait
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("world")
		wg.Done()
	}()

	// mutexes (critical section) - 1 таска в моменте
	mut := &sync.Mutex{}
	mut.Lock()

	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Exit")
}
