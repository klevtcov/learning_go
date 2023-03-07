package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
