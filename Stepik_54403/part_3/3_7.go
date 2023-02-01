package m3_7

import (
	"fmt"
	// "io"
	// "os"
	// "strconv"
	"time"
)

func main() {

	var input string
	// input, err := io.Read(os.Stdin)
	fmt.Scan(&input)
	// input := "1986-04-16T05:20:00+06:00"
	// inputStr := string(input)

	inputTime, err := time.Parse("2006-01-02T15:04:05Z07:00", input)
	if err != nil {
		panic(err)
	}

	fmt.Println(inputTime.Format(time.UnixDate))
}

// time.RFC3339 и time.UnixDate
// На стандартный ввод подается строковое представление даты и времени в следующем формате:

// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

// Wed Apr 16 05:20:00 +0600 1986

// Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.
//
//
// Изначальные значения уже были в формате time.RFC3339б можно было парсить через готовый формат
//
// func main() {
// 	var task string
// 	fmt.Scan(&task)

// 	date, err := time.Parse(time.RFC3339, task)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print(date.Format(time.UnixDate))
// }
//
//
// -------------------------------
//
// func main() {
//     var s string
//     fmt.Scan(&s)
//     if rawdate, err := time.Parse(time.RFC3339, s); err == nil{
//         fmt.Println(rawdate.Format(time.UnixDate))
//     }

// }
