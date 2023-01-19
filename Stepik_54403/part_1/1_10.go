package m1_10

import "fmt"

func main() {

	// Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
	// Квадрат каждого числа должен выводится в новой строке.
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i * i)
	// }

	// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных
	// числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
	// var a, b, sum int
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// for i := a; i <= b; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// без доп. переменной
	// for i:=a+1; i<=b; i++{
	//     a+=i
	// }
	// fmt.Print(a)

	// Через математику - сумма чисел
	// fmt.Println((b - a + 1) * (a + b) / 2)

	// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
	// Программа в первой строке получает на вход число n - количество чисел в последовательности,
	// во второй строке -- n чисел, входящих в данную последовательность.
	// 5
	// 38 24 800 8 16
	// вывод - 40
	// var a, b, sum int
	// fmt.Scan(&a)
	// for i := 1; i <= a; i++ {
	// 	fmt.Scan(&b)
	// 	if 9 < b && b < 100 && b%8 == 0 {
	// 		sum += b
	// 	}
	// }
	// fmt.Println(sum)

	// Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
	// которые равны ее наибольшему элементу.
	// Формат входных данных
	// Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в
	// последовательность не входит, а служит как признак ее окончания).
	// Формат выходных данных 1 3 3 1 0 - 2
	// Выведите ответ на задачу.
	// var n, count, highest int
	// count = 0
	// highest = 0
	// for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
	// 	if n > highest {
	// 		count = 1
	// 		highest = n
	// 	} else if n == highest {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	// Разносим проверку условия и изменения счетчика в разные конструкции
	// for fmt.Scan(&a); a!=0; fmt.Scan(&a){
	//     if a> s {s = a; p = 0}
	//     if a==s {p++}
	// }

	// можно присвоения делать в одну строчку
	// max, count = a, 1

	// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
	// Входные данные
	// Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
	// Выходные данные
	// Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d. Если такого числа нет - выводить ничего не нужно.
	// var n, c, d int
	// fmt.Scan(&n, &c, &d)
	// for i := 0; i <= n; i++ {
	// 	if i%c == 0 && i%d != 0 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	// Начинаем с "C" и далььше итерируемся только по кратным С числам, увеличивая каждый шаг счетчик на С
	// for i:=c; i<=n; i+=c{
	//     if i%d!=0 {
	//         fmt.Print(i)
	//         break
	//     }

	// Напишите программу, которая считывает целые числа с консоли по одному числу в строке.
	// Для каждого введённого числа проверить:
	// если число меньше 10, то пропускаем это число;
	// если число больше 100, то прекращаем считывать числа;
	// в остальных случаях вывести это число обратно на консоль в отдельной строке.
	// var a int
	// for fmt.Scan(&a); a <= 100; fmt.Scan(&a) {
	// 	if a < 10 {
	// 		continue
	// 	} else {
	// 		fmt.Println(a)
	// 	}
	// }
	// альтернатива, без continue
	// if a >=10 {fmt.Println(a)}

	// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
	// Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
	// Входные данные
	// Программа получает на вход три натуральных числа: x, p, y.
	// Выходные данные
	// Программа должна вывести одно целое число.
	// var x, p, y int
	// fmt.Scan(&x, &p, &y)
	// year := 1
	// for {
	// 	x += x * p / 100
	// 	// fmt.Println(x)
	// 	if x > y {
	// 		fmt.Println(year)
	// 		break
	// 	}
	// 	year++
	// }

	// Производим увеличесние депозита в шаге цикла
	// count := 0
	// fmt.Scan(&x,&p,&y)

	// for ; x < y; x += x*p/100 {
	//     count++
	// }

	//  Тут условие if переезжает в условие цикла
	// var depositSize, percent, threshold, numYears int

	// fmt.Scan(&depositSize, &percent, &threshold)

	// for depositSize < threshold {
	// 	depositSize += depositSize * percent / 100
	// 	numYears++
	// }
	// fmt.Println(numYears)

	// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
	// Входные данные 564 8954
	// Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
	// Выходные данные 5 4
	// Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
	// Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

	// решил через строки
	// var a, b string
	// fmt.Scan(&a, &b)
	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < len(b); j++ {
	// 		if a[i] == b[j] {
	// 			fmt.Print(string(a[i]), " ")
	// 		}
	// 	}
	// }
	fmt.Println("End")
}
