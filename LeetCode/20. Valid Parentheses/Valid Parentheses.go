package main

import "fmt"

// func isValid(s string) bool {
// 	dict := map[string]int{
// 		"(": 0,
// 		")": 0,
// 		"[": 0,
// 		"]": 0,
// 		"{": 0,
// 		"}": 0,
// 	}
// 	for _, val := range s {
// 		dict[string(val)] += 1
// 	}

// 	if dict["("] == dict[")"] && dict["["] == dict["]"] && dict["{"] == dict["}"] {
// 		return true
// 	}
// 	return false
// }

func isValid(s string) bool {
	for i := 0; i < len(s); i = i + 2 {
		switch string(s[i]) {
		case "(":
			if string(s[i+1]) != ")" {
				return false
			}
		case "[":
			if string(s[i+1]) != "]" {
				return false
			}
		case "{":
			if string(s[i+1]) != "}" {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()[]{}["))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]")) // false
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Input: s = "()"
// Output: true

// Input: s = "()[]{}"
// Output: true

// Input: s = "(]"
// Output: false
