package m3_7_1

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dateStr := scanner.Text()

	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		panic(err)
	}

	if date.Hour() >= 13 {
		date = date.Add(time.Hour * 24)
	}
	fmt.Println(date.Format("2006-01-02 15:04:05"))

}

// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

// 2020-05-15 08:00:00
// 2020-05-15 13:01:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.

// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
//
//
// Формат можно хранить в константе
//
// const F = "2006-01-02 15:04:05"
// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	input := scanner.Text()
// 	t, err := time.Parse(F, input)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if t.Hour() >= 13 {
// 		t = t.Add(time.Hour * 24)
// 	}
// 	fmt.Println(t.Format(F))
// }
//
//
// -----------------------------
