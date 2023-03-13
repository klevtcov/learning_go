// # Бинарный поиск - сравниваем середину отсортированного списка с требуемым значением,
// # если не попали - сужаем поиск ещё в два раза - делим пополам следующий диапазон и т.д.

package main

// func BinarySearch(arr []int, n int) int {
// 	low := 0
// 	high := len(arr) - 1

// 	for low <= high {
// 		mid := (low + high) / 2
// 		guess := arr[mid]
// 		if guess == n {
// 			return mid
// 		} else if guess > n {
// 			high = mid - 1
// 		} else if guess < n {
// 			low = mid + 1
// 		}
// 	}
// 	return -1
// }
