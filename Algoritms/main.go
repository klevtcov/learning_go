package main

import (
	"fmt"
	// "github.com/klevtcov/algoritms"
)

// Бинарный поиск, принимает отсортированный массив и число, возвращает индекс числа в массиве есть найден и -1 если нет.
func BinarySearch(arr []int, n int) int {
	low := 0
	high := len(arr) - 1

	// Сравниваем середину списка с требуемым значением, если не равно - сужаем диапазон в два раза, сдвигая верхнюю или нижнюю границу
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == n {
			return mid
		} else if guess > n {
			high = mid - 1
		} else if guess < n {
			low = mid + 1
		}
	}
	return -1
}

// Находит наименьшее значение в массиве и возвращает его индекс.
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for ind, val := range arr {
		if val < smallest {
			smallest = val
			smallest_index = ind
		}
	}
	return smallest_index
}

// Сортировка списка через поиск наименьшего/наибольшего значения. Перебираем значения исходного списка и добавлем его в новый, удаляя в исходном.
func SelectionSort(arr []int) []int {
	sortedArr := make([]int, 0, len(arr))
	for range arr {
		smallest := findSmallest(arr)
		sortedArr = append(sortedArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return sortedArr
}

func main() {
	// fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
	// fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7))
	// fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	// fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 12))

	fmt.Println(SelectionSort([]int{5, 7, 3, 4, 2, 9, 1, 8, 6}))
	fmt.Println(SelectionSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 16, 47, -2}))
	fmt.Println(SelectionSort([]int{5, 7, 3, 4, 2, 9, 1, 8, 6}))
}
