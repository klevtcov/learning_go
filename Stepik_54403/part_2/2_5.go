package m2_5

import (
	"fmt"
	"unicode"
)

func main() {

	var str string
	fmt.Scan(&str)
	runeStr := []rune(str)
	if len(runeStr) < 5 {
		fmt.Println("Wrong password")
		return
	} else {
		for _, val := range runeStr {
			if unicode.IsDigit(val) || unicode.Is(unicode.Latin, val) {
				continue
			} else {
				fmt.Println("Wrong password")
				return
			}
		}
	}
	fmt.Println("Ok")

}

//
// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы
// латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям -
// вывести "Ok", иначе вывести "Wrong password"
// проверка символа на цифру
//  fmt.Println(unicode.IsDigit('1')) // true
// например, проверка на латиницу:
// fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false

// можно проверять сразу на два условия
// if unicode.In(ch, unicode.Latin, unicode.Digit) == false {
// 	output = "Wrong password"
// }
//
//

// var str string
// fmt.Scan(&str)
// runeStr := []rune(str)
// if len(runeStr) < 5 {
// 	fmt.Println("Wrong password")
// 	return
// } else {
// 	for _, val := range runeStr {
// 		if unicode.IsDigit(val) || unicode.Is(unicode.Latin, val) {
// 			continue
// 		} else {
// 			fmt.Println("Wrong password")
// 			return
// 		}
// 	}
// }
// fmt.Println("Ok")

//
//
// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
// Sample Input: zaabcbd
// Sample Output: zcd

// var str string
// fmt.Scan(&str)
// runeStr := []rune(str)
// for _, val := range runeStr {
// 	if strings.Count(string(runeStr), string(val)) == 1 {
// 		fmt.Printf(string(val))
// 	}
// }

//
// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

// Sample Input: ihgewlqlkot
// Sample Output: hello

// var str string
// fmt.Scan(&str)
// strRune := []rune(str)
// var result []string
// if len(strRune) < 2 {
// 	fmt.Println(str)
// 	return
// }
// for i := 1; i < len(strRune); i = i + 2 {
// 	result = append(result, string(strRune[i]))
// }

// for _, val := range result {
// 	fmt.Printf(val)
// }

//
//
// Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1

// Sample Input: awesome es
// Sample Output: 2

// var x, s string
// fmt.Scan(&x, &s)
// if strings.Contains(x, s) {
// 	fmt.Println(strings.Index(x, s)) // -1 возвращается по умолчанию.
// } else {
// 	fmt.Println(-1)
// }

//
//
// На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой. Если
// строка правильная - вывести Right иначе - вывести Wrong
// Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы
// можете считать полностью строку за раз с помощью bufio:

// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// runeText := []rune(text)
// if strings.HasSuffix(text, ".") && unicode.IsUpper(runeText[0]) {
// 	fmt.Println("Right")
// } else {
// 	fmt.Println("Wrong")
// }

//  из комментов – Для работы со строками лучше переводить всё в руны из-за Unicode.
// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// runes := []rune(text)
// if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
// 	fmt.Println("Right")
// } else {
// 	fmt.Println("Wrong")
// }

// На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом -
// вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

// func isPalindrome(s string) string {
// 	runeS := []rune(s)
// 	for i := 0; i < len(runeS)/2; i++ {
// 		if runeS[i] != runeS[len(runeS)-i-1] {
// 			return "Нет"
// 		}
// 	}
// 	return "Палиндром"
// }

// можно не брать символ по индексу, а взять сразу его:
// str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// text := []rune(str)
// for i, j := range text{
// 	if text[len(text)-i-1] != j{
// 		fmt.Print("Нет")
// 		return
// 	}
// }
// fmt.Print("Палиндром")

// var str string
// str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// fmt.Println(isPalindrome(str))

// s := "топот"
// a := "закз"
// b := "абба"
// c := "манеенам"
// d := "надда"
// fmt.Println(isPalindrome(s), " - s")
// fmt.Println(isPalindrome(a), " - a")
// fmt.Println(isPalindrome(b), " - b")
// fmt.Println(isPalindrome(c), " - c")
// fmt.Println(isPalindrome(d), " - d")
