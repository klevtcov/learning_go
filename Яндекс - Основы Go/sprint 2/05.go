package ma005

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Почта"`
	DateOfBirth time.Time
}

func main() {
	alex := Person{
		Name:  "Aлекс",
		Email: "alex@yandex.ru",
	}

	data, err := json.Marshal(alex)
	if err != nil {
		fmt.Println("Can't marshal data")
	}

	fmt.Printf("Person %v", string(data)) // Person {"Имя":"Aлекс","Почта":"alex@yandex.ru","DateOfBirth":"0001-01-01T00:00:00Z"}
}

// Для сериализации используется функция json.Marshal() пакета json. Дана структура:
// type Person struct {
//     Name        string
//     Email       string
//     DateOfBirth time.Time
// }
// Напишите код, который будет сериализовывать структуру в json-строку следующего вида:
// {
//   "Имя": "Aлекс",
//   "Почта": "alex@yandex.ru"
// }
