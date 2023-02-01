package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Global []struct {
	Id int `json:"global_id"`
}

// func (g *Global) Summ() int {
// 	summ := 0
// 	for i := 0; i < len(g.I); i++ {
// 		summ += g[i].Id
// 	}
// 	for _, val := range g {
// 		summ += g.Id
// 	}
// 	return summ
// }

func main() {
	file, err := os.Open("./3_6_1.json")
	if err != nil {
		fmt.Println("can't open a file")
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("can't read a file")
		return
	}

	var dataJson Global
	if err := json.Unmarshal(data, &dataJson); err != nil {
		fmt.Println("Can't unmarshal json")
		return
	}

	fmt.Println(dataJson[0])

	// var result int
	// for _, val := range dataJson {
	// 	result += val
	// }
	// for i:=0; i<len(dataJson); i++{
	// }
	// obj.Tests[i].Id
	// }
}

//    "net/http"
// )

// func main(){

//    resp, err := http.Get(urlDownload)
//    if err != nil {
//       return
//    }
//    defer resp.Body.Close()
//    r := csv.NewReader(resp.Body)

// Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем использовать справочник ОКВЭД
// (Общероссийский классификатор видов экономической деятельности), который был размещен на web-портале data.gov.ru.

// Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные, которые
// вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.

// Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать
// сумму полей global_id всех элементов, закодированных в файле.
