package m1_9

import "fmt"

func main() {

	// На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
	// если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

	// var a int
	// fmt.Scan(&a)
	// if a < 0 {
	// 	fmt.Println("Число отрицательное")
	// } else if a > 0 {
	// 	fmt.Println("Число положительное")
	// } else {
	// 	fmt.Println("Ноль")
	// }

	// По данному трехзначному числу определите, все ли его цифры различны.
	// Формат входных данных
	// На вход подается одно натуральное трехзначное число.
	// Формат выходных данных
	// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

	// var a int
	// fmt.Scan(&a)
	// a1 := a / 100
	// a2 := a / 10 % 10
	// a3 := a % 10
	// if a1 != a2 && a2 != a3 && a1 != a3 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)

	// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
	// Формат входных данных
	// На вход дается натуральное число, не превосходящее 10000.
	// Формат выходных данных
	// Выведите одно целое число - первую цифру заданного числа.
	// var a int
	// fmt.Scan(&a)
	// switch {
	// case 0 <= a && a <= 9:
	// 	fmt.Println(a)
	// case 10 <= a && a <= 99:
	// 	fmt.Println(a / 10)
	// case 100 <= a && a <= 999:
	// 	fmt.Println(a / 100)
	// case 1000 <= a && a <= 9999:
	// 	fmt.Println(a / 1000)
	// case a == 10000:
	// 	fmt.Println(1)
	// }

	// из комментов
	// switch {
	// case a/10000 >= 1:
	// 	fmt.Println(a / 10000)
	// case a/1000 >= 1:
	// 	fmt.Println(a / 1000)
	// case a/100 >= 1:
	// 	fmt.Println(a / 100)
	// case a/10 >= 1:
	// 	fmt.Println(a / 10)
	// default:
	// 	fmt.Println(a)
	// }

	// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма
	// первых трёх цифр совпадает с суммой трёх последних.
	// Формат входных данных
	// На вход подается номер билета - одно шестизначное  число.
	// Формат выходных данных
	// Выведите "YES", если билет счастливый, в противном случае - "NO".
	// Sample Input: 613244

	// var a int
	// fmt.Scan(&a)
	// // a = 613245
	// a1 := a / 100000
	// a2 := (a % 100000) / 10000
	// a3 := (a % 10000) / 1000
	// a4 := (a % 1000) / 100
	// a5 := (a % 100) / 10
	// a6 := a % 10
	// // fmt.Println(fmt.Sprintf("%T", a1))
	// if a1+a2+a3 == a4+a5+a6 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	// // if a1 != a2 && a2 != a3 && a1 != a3 {
	// // 	fmt.Println("YES")
	// // } else {
	// // 	fmt.Println("NO")
	// // }

	// fmt.Println(a1, a2, a3, a4, a5, a6)

	// решения из комментов
	// var a,b,c,d,e,f int
	// fmt.Scanf("%1d%1d%1d%1d%1d%1d",&a,&b,&c,&d,&e,&f)

	// if ((d+e+f) == (a+b+c)){
	//     fmt.Print("YES")
	// }else{
	//     fmt.Println("NO")
	// }

	// Требуется определить, является ли данный год високосным, напомним:
	// Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
	// - кратен 400;
	// - кратен 4, но не кратен 100.
	// Входные данные
	// Вводится единственное число - номер года (целое, положительное, не превышает 10000).
	// Выходные данные
	// Требуется вывести слово YES, если год является високосным и NO - в противном случае.

	// var a int
	// fmt.Scan(&a)

	// if (a%400 == 0) || (a%4 == 0) && (a%100 != 0) {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	fmt.Println("end")
}
