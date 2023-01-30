package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "unicode"
)

func main() {

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// s := scanner.Text()
	// input = strings.TrimRight(input, "\n\r")
	// input, _ = bufio.NewReader(os.Stdin).ReadString('\r')

	str, err := bufio.NewReader(os.Stdin).ReadString('\r')
	if err != nil && err != io.EOF {
		fmt.Println("error: ", err)
	}
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, ",", ".", -1)
	var strArr []string
	strArr = strings.Split(str, ";")
	var str1, str2 string
	str1 = strArr[0]
	str2 = strArr[1]
	// var strFloat1, strFloat2 float64
	strFloat1, _ := strconv.ParseFloat(string(str1), 64)
	strFloat2, _ := strconv.ParseFloat(string(str2), 64)
	fmt.Println(strFloat1)
	fmt.Println(strFloat2)

	// for ind, val := range strArr {
	// 	fmt.Println(ind, " ", val, ".")
	// }
	// var strArr1 string
	// strArr1 = string(strArr1[1])
	// fmt.Println(strArr1)
	// fmt.Println(strArr)
	// fmt.Println(strArr[0])
	// fmt.Println(strArr[1])
	// str1, _ := strconv.ParseFloat(string(strArr[0]), 64)
	// str2, _ := strconv.ParseFloat(string(strArr[1]), 64)

	// fmt.Printf("%T ", str1)
	// fmt.Println(str1)
	// fmt.Printf("%T ", str2)
	// fmt.Println(str2)
	// result := str1 / str2
	// fmt.Printf("%.4f", result)

	// fmt.Println(" ")
}

// Sample Input: 1 149,6088607594936;1 179,0666666666666
// Sample Output: 0.9750
// Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv,
// или даже bufio - вы не ограничены в выборе способа решения задачи. В решениях мы поделимся своими способами решения
// этой задачи, предлагаем вам сделать то же самое.

// В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в
// виде пробела, кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в
// формат CSV, где в качестве разделителя используется символ ";".

// На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления
// первого числа на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

// P.S. небольшое отступление, связанное с чтением из стандартного ввода. Кто-то захочет использовать для этого пакет
// bufio.Reader. Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете
// ошибку EOF, а точнее (io.EOF - End Of File). На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом)
// прочитан до конца. Чтобы ошибка была обработана правильно, вы можете поступить так:
// if err != nil && err != io.EOF {
// 	...
// }
//

//
//
// Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то
// логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа
// string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше. Поэтому он решил
// подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние
// символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они
// уже импортированы, вам ничего импортировать не нужно!
// Считывать и выводить ничего не нужно!
// Ваша функция должна называться adding() !
// Считайте что функция и пакет main уже объявлены!
//
// func adding(str1, str2 string) int64 {
// 	return clearStringToInteger(str1) + clearStringToInteger(str2)
// }

// func clearStringToInteger(str string) int64 {
// 	var result []string
// 	for _, val := range str {
// 		if unicode.IsDigit(val) {
// 			result = append(result, string(val))
// 		}
// 	}
// 	resultStr := strings.Join(result, "")
// 	resultInt, _ := strconv.Atoi(resultStr)
// 	return int64(resultInt)
// }
//
// Магия через Трим и вложенные функции
// func help(r rune) bool {
//     return !unicode.IsDigit(r)
// }

// func adding(a, b string) int64 {
//     n, _ := strconv.ParseInt(strings.TrimFunc(a, help), 10, 64)
//     m, _ := strconv.ParseInt(strings.TrimFunc(b, help), 10, 64)
//     return n + m
// }
//
//
//
// Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
// Считывать и выводить ничего не нужно!
// Считайте что функция main и пакет main уже объявлены!
// func convert(n int64) uint16 {
// 	return uint16(n)
// }

// именованный ретёрн
// func convert(a int64) (b uint16) {
//     return
// }
