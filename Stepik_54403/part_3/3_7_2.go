package m3_7_2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const F = "02.01.2006 15:04:05"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strScan := scanner.Text()
	var date []string
	date = strings.Split(strScan, ",")
	date1, _ := time.Parse(F, date[0])
	date2, _ := time.Parse(F, date[1])

	if date1.Before(date2) {
		date1, date2 = date2, date1
	}
	fmt.Println(date1.Sub(date2))

}

// 11.03.2018 14:00:15,12.03.2018 14:00:15
// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
// 13.03.2018 14:00:15,12.03.2018 14:00:15
// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
// 24h0m0s
//
//
//
// Можно просто привести разницу во времени к положительному значению. и чтение данных через CSV
//
// func main() {
// 	row, _ := csv.NewReader(os.Stdin).Read()
// 	x, _ := time.Parse("02.01.2006 15:04:05", row[0])
// 	y, _ := time.Parse("02.01.2006 15:04:05", row[1])
// 	diff := x.Sub(y)
// 	if diff < 0 {
// 		diff *= -1
// 	}
// 	fmt.Println(diff)
// }
//
// ------------------------
//
// Взятие модуля из библиотеки math
//
// func main() {
// 	var DATE_FORMAT = "02.01.2006 15:04:05"

// 	dates, _ := csv.NewReader(os.Stdin).Read()
// 	date_1, _ := time.Parse(DATE_FORMAT, dates[0])
// 	date_2, _ := time.Parse(DATE_FORMAT, dates[1])

// 	fmt.Println(time.Duration(math.Abs(float64(date_1.Sub(date_2)))).Round(time.Second))
// }
// --------------------------------------
//
// Обработку стриоки и приведение её к дате вынесли в отдельную функцию
//
// func _getTime() func() time.Time {
// 	reader := bufio.NewReader(os.Stdin)
// 	const layout string = "02.01.2006 15:04:05"

// 	return func() time.Time {
// 		timeInString, _ := reader.ReadString(',')
// 		timeInString = strings.Trim(timeInString, "\n, ")
// 		t, _ := time.Parse(layout, timeInString)
// 		return t
// 	}
// }

// func main() {
// 	getTime := _getTime()

// 	lastTime, firstTime := getTime(), getTime()

// 	if firstTime.After(lastTime) {
// 		lastTime, firstTime = firstTime, lastTime
// 	}

// 	difference := lastTime.Sub(firstTime)

// 	fmt.Println(difference)
// }
