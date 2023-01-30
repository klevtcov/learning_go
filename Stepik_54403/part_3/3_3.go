package m3_3

import (
	"fmt"
	"strconv"
	// "strings"
	// "unicode/utf8"
)

func main() {

	fn := func(i uint) uint {
		str := strconv.Itoa(int(i))
		runes := []rune(str)
		fmt.Println(runes)
		var result []rune
		for _, val := range runes {
			val -= '0'
			if val%2 == 0 && val != 0 {
				result = append(result, val+'0')
			}
		}
		var s string
		s = string(result)
		si, _ := strconv.Atoi(s)
		if si == 0 {
			return uint(100)
		}
		return uint(si)
	}
	fmt.Println(fn(6554))

}

//
// Проще через математику решать, чем через преобразование типов. так говнокода больше, но интеерснее, и ближе к теме
//
// fn := func(i uint) uint {
// 	str := strconv.Itoa(int(i))
// 	runes := []rune(str)
// 	fmt.Println(runes)
// 	var result []rune
// 	for _, val := range runes {
// 		val -= '0'
// 		if val%2 == 0 && val != 0 {
// 			result = append(result, val+'0')
// 		}
// 	}
// 	var s string
// 	s = string(result)
// 	si, _ := strconv.Atoi(s)
// 	if si == 0 {
// 		return uint(100)
// 	}
// 	return uint(si)
// }
//
// Решения здорового человека:
//
// fn := func(n uint) uint {
// 	str := strconv.FormatUint(uint64(n), 10)
// 	apstr := []rune{}
// 	for _, elem := range str {

// 		if elem%2 == 0 && elem != '0' {
// 			apstr = append(apstr, elem)
// 		}
// 	}

// 	bar, err := strconv.ParseUint(string(apstr), 10, 64)
// 	if bar == 0 {
// 		return 100
// 	}
// 	if err != nil {
// 		panic(err)
// 	}

// 	return uint(bar)
// }
//
//  Через математику:
//
// fn := func (x uint) (y uint) {
//     for k := uint(1); x > 0; x /= 10 {
//         if d := x % 10; d % 2 == 0 && d != 0 {
//             y += k * d
//             k *= 10
//         }
//     }
//     if y == 0 {
//         y = 100
//     }
//     return
// }
//
//
// fn := func(X uint) uint {
//     var x uint
//     s := strconv.FormatUint(uint64(X), 10)
//     for i := range s {
//         if s[i] % 2 == 0 && s[i] != '0' {
//             x = x * 10 + uint(s[i] - '0')
//         }
//     }
//     if x == 0 {
//         x = 100
//     }
//     return x
// }
//
//
//
//
// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main
// в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные
// цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
//
//

//
// пример работы функции - можем вызвать функцию и передавть в неё две другие и два аргумента. логика прописана внутри функции
// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	fmt.Println(sum(up10, del2, a, b))
// }
// func sum(fa func(int) int, fb func(int) int, a, b int) int {
// 	return fa(a) + fb(b)
// }
// func up10(x int) int {
// 	return x * 10
// }
// func del2(c int) int {
// 	return c / 2
// // }
//
//
// x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
// fmt.Printf("%T", x)
