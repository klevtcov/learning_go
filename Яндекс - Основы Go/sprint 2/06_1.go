package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func squareArea(x float64) float64 {
	return x * x
}

func circleArea(x float64) float64 {
	return 3.14 * x * x
}

func triangleArea(x float64) float64 {
	return math.Sqrt(3) * x * x / 4
}

func area(f figures) (func(float64) float64, bool) {
	//Функция должна возвращать:
	// функцию для вычисления площади фигуры;
	// true, если фигура известна;
	// false, если фигура неизвестна.
	switch f {
	case square:
		return squareArea, true
	case circle:
		return circleArea, true
	case triangle:
		return triangleArea, true

	default:
		return nil, false
	}
}

// func area(f figures) (func(float64) float64, bool) {
//     switch f {
//     case square:
//         return func(x float64) float64 { return x * x }, true
//     case circle:
//         return func(x float64) float64 { return 3.142 * x * x }, true
//     case triangle:
//         return func(x float64) float64 { return 0.433 * x * x }, true
//     default:
//         return nil, false
//     }
// }

func main() {
	// myFigure := square
	// myFigure := circle
	myFigure := triangle
	// myFigure := octaedr

	var x float64
	x = 7

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(x)
	fmt.Printf("Площадь %v равна %.2f", myFigure, myArea)

}
