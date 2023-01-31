package m_task2

import (
	"bufio"
	// "encoding/csv"
	"fmt"
	"io"
	"os"
	// "time"
)

// полезный туториал
// https://metanit.com/go/tutorial/8.9.php

func main() {

	// чатаем файл, ошибки игнорируем
	file, _ := os.Open("./task.data")

	// отложенное закрытие файлы, выполнится в конце
	defer file.Close()

	// задаём счетчик и передаём в читалку файл
	counter := 1
	reader := bufio.NewReader(file)

	// итерируемся по ридеру, ориентируясь на ;. если строка равна нулю - выводим значение счетчика, если нет - увеличиваем счетчик на 1 и идём дальше
	// если ошибкуа - завершаем цикл
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if line == "0;" {
			fmt.Print(counter)
			return
		}
		counter++
	}
}

//
// Другие решения
//
//Через массив
// func main() {
// 	file, _ := os.ReadFile("task.data")
// 	arr := strings.Split(string(file), ";")
// 	for i := range arr {
// 	   if arr[i] == "0" {
// 		  fmt.Println(i + 1)
// 	   }
// 	}
//  }
// ----------------------------
//
// Через CSV
//
// func main() {

// 	file, err := os.Open("task.data")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	reader := csv.NewReader(file)
// 	reader.Comma = ';'
// 	record, err := reader.Read()
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i, s := range record {
// 		if s == "0" {
// 			fmt.Printf("Индекс %v число %v\n", i+1, s)
// 		}
// 	}
// }
