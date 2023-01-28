package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(adding("hjj_1kjjd0", "hsadsad3kklasdjns5asd"))
	// fmt.Println(" ")
}

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
