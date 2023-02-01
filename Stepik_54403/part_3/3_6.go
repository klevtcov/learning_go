package m3_6

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Groups struct {
	ID       int    `json:"ID"`
	Number   string `json:"Number"`
	Year     int    `json:"Year"`
	Students []struct {
		LastName   string `json:"LastName"`
		FirstName  string `json:"FirstName"`
		MiddleName string `json:"MiddleName"`
		Birthday   string `json:"Birthday"`
		Address    string `json:"Address"`
		Phone      string `json:"Phone"`
		Rating     []int  `json:"Rating"`
	} `json:"Students"`
}

type Average struct {
	Average float32
}

func main() {

	// file, err := os.Open("./3_6.json")
	// if err != nil {
	// 	fmt.Println("can't open a file")
	// 	return
	// }
	// defer file.Close()
	// data, err := io.ReadAll(file)

	data, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("can't read a file")
		return
	}

	var dataJson Groups
	if err := json.Unmarshal(data, &dataJson); err != nil {
		fmt.Println("Can't unmarshal json")
		return
	}

	students := len(dataJson.Students)
	var score int
	for i := 0; i < len(dataJson.Students); i++ {
		score += len(dataJson.Students[i].Rating)
	}

	result := Average{
		Average: float32(float32(score) / float32(students)),
	}

	resultJson, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = os.Stdout.Write(resultJson)

}

// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
//         {
//             // ...
//         }
//     ]
// }
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
// {
//     "Average": 14.1
// }
// Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).
// Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для вас: в уроках об интерфейсах и работе с файлами
// мы рассказывали, что стандартный ввод / вывод - это файлы, к которым можно обратиться через os.Stdin и os.Stdout соответственно,
// они удовлетворяют интерфейсам io.Reader и io.Writer, из них можно читать, в них можно писать.
// Один из способов чтения, использовать ioutil.ReadAll:
// data, err := ioutil.ReadAll(os.Stdin)
// if err != nil {
//     // ...
// }
// data - тип []byte
//
//
// Вариант решения с взятием только нужных полей из JSON
//
// type (
// 	Student struct {
// 		Rating []int
// 	}
// 	Group struct {
// 		Students []Student
// 	}
// )

// func main() {
// 	students := new(Group)

// 	bytes, _ := ioutil.ReadAll(os.Stdin)
// 	_ = json.Unmarshal(bytes, students)

// 	var cnt int
// 	for _, student := range students.Students {
// 		cnt += len(student.Rating)
// 	}

// 	average := float64(cnt) / float64(len(students.Students))

// 	bytes, _ = json.MarshalIndent(map[string]float64{"Average": average}, "", "    ")

// 	fmt.Printf("%s", bytes)
// }
//
// --------------------------------
//
// С объявлением анонимной структуры для вывода
//
// import (
//     "encoding/json"
//     "os"
// )

// type Group struct {
// 	Students []struct {
// 		Rating []int
// 	}
// }
// func main() {
//     var group Group
//     if err := json.NewDecoder(os.Stdin).Decode(&group); err != nil {
//         panic(err)
//     }
//     var avg float64
//     for _, st := range group.Students {
//         avg += float64(len(st.Rating))
//     }
//     avg /= float64(len(group.Students))
//     e := json.NewEncoder(os.Stdout)
//     e.SetIndent("", "    ")
//     e.Encode(struct {Average float64}{avg})
// }
//------------------------------------
//
// Через метод структуры
//
// type Group struct {
// 	ID       int       `json:"ID"`
// 	Number   string    `json:"Number"`
// 	Year     int       `json:"Year"`
// 	Students []struct {
// 	LastName   string `json:"LastName"`
// 	FirstName  string `json:"FirstName"`
// 	MiddleName string `json:"MiddleName"`
// 	Birthday   string `json:"Birthday"`
// 	Address    string `json:"Address"`
// 	Phone      string `json:"Phone"`
// 	Raiting    []int  `json:"Raiting"`
//     } `json:"Students"`
// }

// func (g *Group)Average()float64{
//     studentCount := len(g.Students)
// 	ratingCount  := 0
// 	for _, student := range g.Students{
// 		ratingCount += len(student.Raiting)
// 	}
//     return float64(ratingCount)/float64(studentCount)
// }

// type ResultStruct struct {
// 	Average float64 `json:"Average"`
// }
// func main() {
// 	var group Group
// 	var result ResultStruct
// 	decoder := json.NewDecoder(os.Stdin)
// 	decoder.Decode(&group)
//     result.Average = group.Average()
// 	buffRes, _ := json.MarshalIndent(&result, "", "    ")
// 	os.Stdout.Write(buffRes)
// }
//
//
//
