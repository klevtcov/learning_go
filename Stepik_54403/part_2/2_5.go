package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
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

}
