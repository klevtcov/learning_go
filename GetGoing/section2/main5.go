package main

import "fmt"

func main() {

	fmt.Println("Hello world")
	// if, else, for, switch case, break continue

	// f := true
	// flag := &f

	// if flag == nil {
	// 	fmt.Println("Value is nil")
	// } else if *flag {
	// 	fmt.Println("True")
	// } else {
	// 	fmt.Println("False")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for {
	// 	fmt.Println("Infinity loop")
	// }

	arr := []string{"Sergey", "Klevtsov", "sk"}

	//  первое счетчик, второе значение
	for i, value := range arr {
		// fmt.Println(i) - 0 1 2
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Sergey"
	mymap["age"] = 39

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}

}
