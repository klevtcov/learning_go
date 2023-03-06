package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums[i+1:] {
			fmt.Println("i: ", i, " v: ", v, " j: ", j, " v2: ", v2)
			if v+v2 == target {
				return []int{i, j + 1}
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
