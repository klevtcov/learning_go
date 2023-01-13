package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {

	fmt.Println("")
	Anything(2.44)
	Anything("Name")
	Anything(struct{}{})

	mymap := make(map[string]interface{})
	// ключ всегда string, но значение любое, т.к. интерфейс
	mymap["name"] = "Sergey"
	mymap["age"] = 39
	fmt.Println(mymap)
	// map[age:39 name:Sergey]
}
