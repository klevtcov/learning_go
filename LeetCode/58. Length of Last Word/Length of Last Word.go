package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.

// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

// func lengthOfLastWord(s string) int {
//     str := strings.Fields(s) //tokenize
//     return len(str[len(str)-1])
// }
