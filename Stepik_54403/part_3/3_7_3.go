package m3_7_3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strScan := scanner.Text()

	strScan = strings.Replace(strScan, " мин. ", "m", 1)
	strScan = strings.Replace(strScan, " сек.", "s", 1)
	dura, _ := time.ParseDuration(strScan)

	unixTime := time.Unix(now, 0).UTC()
	unixTime = unixTime.Add(dura)

	fmt.Println(unixTime.Format(time.UnixDate))
}

// - replace-ом заменил " мин. " и " сек." на формат удобный для ParseDuration
// - к Unix времени добавил (time.Add) полученный duration, ну и Format(time.UnixDate) для вывода

// func ParseDuration(s string) (Duration, error)
// преобразует строку в Duration с использованием аннотаций:
// "ns" - наносекунды,
// "us" - микросекунды,
// "ms" - миллисекунды,
// "s" - секунды,
// "m" - минуты,
// "h" - часы.
// dur, err := time.ParseDuration("1h12m3s")
// if err != nil {
// 	panic(err)
// }
// fmt.Println(dur.Round(time.Hour).Hours()) // 1

// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в
// формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести
// (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.
// Sample Input: 12 мин. 13 сек.
// Sample Output: Fri May 15 19:28:18 UTC 2020
//
//
// Считываем сразу в две переменные и внутри строки добавляем их как минуты и секунды. никаких парсингов в дюрейшн
//
// func main() {
// 	var m, s time.Duration
// 	fmt.Scanf("%d мин. %d сек.", &m, &s)
// 	t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
// 	fmt.Println(t.Format(time.UnixDate))
// }
//
// ---------------------------
//
// Почти похоже, но т.к. Unix это у нас секунды от 1970 года, то можем сразу запихнуть секунды в конвертацию
//
// func main() {
// 	var m, s int64
// 	fmt.Scanf("%d мин. %d сек.", &m, &s)
// 	t := time.Unix(1589570165 + m*60 + s, 0).UTC()
// 	fmt.Println(t.Format(time.UnixDate))
// }
//---------------------
//
// Множественные замены, сразу под кейсы с часами подходит
//
// func main() {
// 	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 	r := strings.NewReplacer("ч.", "h", "мин.", "m", "сек.", "s", " ", "", "\n", "")
// 	dur, _ := time.ParseDuration(r.Replace(input))
// 	fmt.Println(time.Unix(1589570165, 0).Add(dur).UTC().Format(time.UnixDate))
// }
// ----------------------------
//
// Скан читает до пробелов, можем сразу считать в нужные переменные
//
// func main() {
//     var min, s1, sec, s2 string
//     fmt.Scan(&min, &s1, &sec, &s2)

//     duration, err := time.ParseDuration(fmt.Sprintf("%sm%ss", min, sec))
//     if err != nil {
//         panic(err)
//     }

//     fmt.Print(time.Unix(now, 0).UTC().Add(duration).Format(time.UnixDate))
// }
// -----------------------
