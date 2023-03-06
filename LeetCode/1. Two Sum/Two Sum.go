package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if v+v2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	// nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Input: nums = [3,3], target = 6
// Output: [0,1]

//
// Можно ускорить проходом с двух сторон
// func twoSum(nums []int, target int) []int {
//     if len(nums) <= 1 {
//         return []int{}
//     }

//     // Time: O(n)
//     var ptr1, ptr2 int = 0, len(nums)-1
//     for ptr1 < ptr2 {
//         current := nums[ptr2] + nums[ptr1]
// 		// Time: O(1) best O(3) worst so O(1)
//         if current == target {
//             return []int{ptr1, ptr2}
//         } else if current > target {
//             ptr2--
//         } else /* if current < target */ {
//             ptr1++
//         }
//     }

//     return []int{}
// }
//
//
//
// Вариант через Хеш-таблицу, но непонятно, в какой момент она создаётся
// Time: O(n)
// Space: O(n)
// func twoSum(nums []int, target int) []int {
//     // Space: O(n)
//     s := make(map[int]int)

//     // Time: O(n)
//     for idx, num := range nums {
//         // Time: O(1)
//         if pos, ok := s[target-num]; ok {
//             return []int{pos, idx}
//         }
//         s[num] = idx
//     }
//     return []int{}
// }
