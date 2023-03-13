// # Бинарный поиск - сравниваем середину отсортированного списка с требуемым значением,
// # если не попали - сужаем поиск ещё в два раза - делим пополам следующий диапазон и т.д.

// # low = 0
// # high = len(list) - 1

// # mid = (low + high) / 2
// # guess = list[mid]
// # if guess < item:
// #     low = mid-1

package main

import "fmt"

func binarySearch(arr []int, n int) int {
	low := 0
	high := len(arr) - 1

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

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 12))
}
