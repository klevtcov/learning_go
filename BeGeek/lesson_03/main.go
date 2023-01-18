package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
)

func main() {
	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// _, i - порядковый номер, значение. Если подчеркивание - то его можно не использовать в дальнейшем
	text := []string{"be", "geek", "hello"}
	for _, i := range text {
		fmt.Println(i)
	}

	// бесконечный цикл
	// for {
	// 	fmt.Println("geek")
	// }

	// прерывание цикла
	for _, i := range text {
		fmt.Println(i)
		if i == "geek" {
			// continue
			break
		}
	}

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {

	// 	fmt.Println(file.Name())
	// }
}
