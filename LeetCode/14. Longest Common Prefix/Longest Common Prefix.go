package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	result := ""
	lowest := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < lowest {
			lowest = len(strs[i])
		}
	}

	for i := 0; i < lowest; i++ {
		pref := strs[0][:i+1]
		flag := true

		for _, val := range strs {
			if strings.HasPrefix(val, pref) {
				continue
			} else {
				flag = false
			}
		}

		if flag {
			result = pref
		} else {
			break
		}

	}
	return result
}

// fmt.Println(strings.HasPrefix("my string", "prefix"))  // false
// fmt.Println(strings.HasPrefix("my string", "my"))      // true

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "floppy"}))
}

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

//

// func longestCommonPrefix(strs []string) string {
//     for i := 0 ;; i++  {
//         for _, str := range strs {
//             if i == len(str) || str[i] != strs[0][i] {
//                 return strs[0][:i]
//             }
//         }
//     }
// }
