package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {

	fmt.Println()
}

// Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект, удовлетворяющий интерфейсу
// fmt.Stringer. Назовем его "Батарейка".

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный). Ваша задача
// считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами на первом этапе типа:
// надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

// В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, но она присутствует.
// Перед этой скобкой присутствует функция (которая вам тоже не видна), принимающая в качестве аргумента объект типа
// fmt.Stringer - batteryForTest, и направляющая его на стандартный вывод, поэтому вам не требуется выводить что-то на печать самостоятельно.
//
//
//
// ------------------------------------------------------------------------------------------------------
//
// Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(), которая возвращает 3 значения
// типа пустой интерфейс. Эта функция использует пакеты encoding/json, fmt, и os - не удаляйте их из импорта. Скорее всего,
// вам понадобится пакет "fmt", но вы можете использовать любой другой пакет для записи в стандартный вывод не удаляя fmt.

// Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64. Третье значение
// в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная математическая операция).
// Но такие идеальные случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем значении
// может не относится к одной из указанных математических операций.

// Результат выполнения программы должен быть таков:

// в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции с точностью до 4
// цифры после запятой (fmt.Printf(%.4f));
// если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке вида: value=полученное_значение:
// тип_значения (например: value=true: bool)
// если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям, сообщение об
// ошибке должно иметь вид: неизвестная операция
// Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения увидели, что оно
// содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных аргументов значения уже не имеет,
// а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.
//
//
// value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
// все полученные значения имеют тип пустого интерфейса
// var value1 interface{} = 12.54
// var value2 interface{} = 31.76
// var operation interface{} = "+"
// if _, ok := value1.(float64); !ok {
// 	fmt.Printf("value=%v: %T", value1, value1)
// 	return
// }
// if _, ok := value2.(float64); !ok {
// 	fmt.Printf("value=%v: %T", value2, value2)
// 	return
// }
// switch {
// case operation == "+":
// 	result := value1.(float64) + value2.(float64)
// 	fmt.Printf("%.4f", result)
// case operation == "-":
// 	result := value1.(float64) - value2.(float64)
// 	fmt.Printf("%.4f", result)
// case operation == "/":
// 	result := value1.(float64) / value2.(float64)
// 	fmt.Printf("%.4f", result)
// case operation == "*":
// 	result := value1.(float64) * value2.(float64)
// 	fmt.Printf("%.4f", result)
// default:
// 	fmt.Println("неизвестная операция")
// }
//
//
// func CheckFloat(a ...interface{}) bool {
// 	for _, k := range a {
// 		if _, ok := k.(float64); !ok { fmt.Printf("value=%v: %T", k, k); return false }
// 	}
// 	return true
// }

// func main() {
// 	value1, value2, operation := readTask()
// 	if CheckFloat(value1, value2) {
// 		switch operation.(string) {
// 		case "+": fmt.Printf("%.4f", value1.(float64) + value2.(float64))
// 		case "-": fmt.Printf("%.4f", value1.(float64) - value2.(float64))
// 		case "*": fmt.Printf("%.4f", value1.(float64) * value2.(float64))
// 		case "/": fmt.Printf("%.4f", value1.(float64) / value2.(float64))
// 		default: fmt.Print("неизвестная операция")
// 		}
// 	}
// }
//
//
// Через интерфейсы по полной
// package main

// import (
// 	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
// 	"fmt"           // пакет используется для проверки ответа, не удаляйте его
// 	"os"            // пакет используется для проверки ответа, не удаляйте его
// )

// type calculator interface {
// 	add() float64
// 	subtract() float64
// 	multi() float64
// 	div() float64
// }

// type Expression struct {
// 	first     float64
// 	second    float64
// 	operation string
// }

// func newExpression(v1, v2, op interface{}) *Expression {
// 	value1 := getFloat(v1)
// 	value2 := getFloat(v2)
// 	operation := getString(op)
// 	return &Expression{value1, value2, operation}
// }

// func getFloat(v interface{}) float64 {
// 	if res, ok := v.(float64); !ok {
// 		panic(fmt.Sprintf("value=%v: %T", v, v))
// 	} else {
// 		return res
// 	}
// }

// func getString(s interface{}) string {
// 	if res, ok := s.(string); !ok {
// 		panic(fmt.Sprintf("value=%v: %T", s, s))
// 	} else {
// 		return res
// 	}
// }

// func (e *Expression) doOperation() float64 {
// 	switch e.operation {
// 	case "+":
// 		return e.add()
// 	case "-":
// 		return e.subtract()
// 	case "*":
// 		return e.multi()
// 	case "/":
// 		return e.div()
// 	default:
// 		panic("неизвестная операция")
// 	}
// }

// func (e *Expression) add() float64 {
// 	return e.first + e.second
// }

// func (e *Expression) subtract() float64 {
// 	return e.first - e.second
// }

// func (e *Expression) multi() float64 {
// 	return e.first * e.second
// }

// func (e *Expression) div() float64 {
// 	return e.first / e.second
// }

// func main() {
// 	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
// 	// все полученные значения имеют тип пустого интерфейса

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	exp := newExpression(value1, value2, operation)
// 	res := exp.doOperation()
// 	fmt.Printf("%0.4f\n", res)
// }
// ------------------------------------------------------------------------------
//
