// # Selection sort – сортировка выбором

// # Массивы - чтение O(1), запись O(n)
// # Списки - чтение O(n), запись O(1)

// # Перебираем список, находим наименьший/наибольший элемент, добавляем его в новый список, удаляя из изначального
// # Повторяем

package main

// func findSmallest(arr []int) int {
// 	smallest := arr[0]
// 	smallest_index := 0
// 	for ind, val := range arr {
// 		if val < smallest {
// 			smallest = val
// 			smallest_index = ind
// 		}
// 	}
// 	return smallest_index
// }

// func SelectionSort(arr []int) []int {
// 	sortedArr := make([]int, len(arr))
// 	for range arr {
// 		smallest := findSmallest(arr)
// 		sortedArr = append(sortedArr, arr[smallest])
// 		arr = append(arr[:smallest], arr[smallest+1:]...)
// 	}
// 	return sortedArr
// }
