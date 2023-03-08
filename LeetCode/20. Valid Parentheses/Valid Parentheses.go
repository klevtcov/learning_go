package main

import "fmt"

func isValid(s string) bool {
	result := make([]string, 0, len(s))
	if len(s)%2 == 1 {
		return false
	}

	for _, val := range s {
		switch string(val) {
		case "(", "[", "{":
			result = append(result, string(val))
		case ")":
			if len(result) < 1 {
				return false
			}
			if result[len(result)-1] != "(" {
				return false
			} else {
				result = result[:len(result)-1]
			}
		case "]":
			if len(result) < 1 {
				return false
			}
			if result[len(result)-1] != "[" {
				return false
			} else {
				result = result[:len(result)-1]
			}
		case "}":
			if len(result) < 1 {
				return false
			}
			if result[len(result)-1] != "{" {
				return false
			} else {
				result = result[:len(result)-1]
			}
		}
	}

	return len(result) <= 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([])[{}]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]")) // false
	fmt.Println(isValid("()[]{}["))
	fmt.Println(isValid("})[]{}"))
	fmt.Println(isValid("(){}}{"))
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

// func isValid(s string) bool {
//     stack := []byte{}

//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' || s[i] == '[' || s[i] == '{' {
//             stack = append(stack, s[i])
//         } else if len(stack) == 0 {
//             return false
//         } else {
//             prev := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             if s[i] == ')' && prev != '(' {
//                 return false
//             } else if s[i] == ']' && prev != '[' {
//                 return false
//             } else if s[i] == '}' && prev != '{' {
//                 return false
//             }
//         }
//     }
//     return len(stack) == 0
// }
//
//
//
// func isValid(s string) bool {
//     if len(s)%2 != 0 {
//         return false
//     }

//     stack := make([]rune,0,len(s))
//     hashMap := map[rune]rune {
//         '}' : '{',
//         ')' : '(',
//         ']' : '[',
//     }
//     for _, char := range s {
//         if _, ok := hashMap[char]; !ok {
//             // this bracket is open
//             stack = append(stack, char)
//         } else {
//             // this bracket is close
//             if len(stack) == 0 || (stack[len(stack)-1] != hashMap[char]) {
//                 return false
//             }
//             stack = stack[:len(stack)-1]
//         }
//     }
//     if len(stack) != 0 {
//         return false
//     }
//     return true
// }
