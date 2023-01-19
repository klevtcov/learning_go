package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Markets []struct {
		Name  string
		Price int
	}
}

// можно передавать структуры как тип данных
// type info struct {
// 	Markets []struct {
// 		Name  status
// 		Price int
// 	}
// }

type status struct {
	Name string
}

func main() {
	text := `{"markets":[{"name":"1", "price":100}, {"name":"2", "price":101}, {"name":"3", "price":102}]}`
	var infos info
	// парсим json - передаём значение text в виде байтов и кладём его в переменную infos
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos)                 // {[{1 100} {2 101} {3 102}]}
	fmt.Println(infos.Markets)         // [{1 100} {2 101} {3 102}]
	fmt.Println(infos.Markets[0].Name) // 1

	var st = status{Name: "Sergey"}
	// st := new(status)
	fmt.Println(st)      // {Sergey}
	fmt.Println(st.Name) // Sergey

	for i := range infos.Markets {
		fmt.Println(i, infos.Markets[i].Name, infos.Markets[i].Price)
	} // 0 1 100  \n   1 2 101   \n    2 3 102

}
