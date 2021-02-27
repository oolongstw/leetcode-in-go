package main

import (
	"fmt"
	"strings"
)

// 解题思路: 左右指针遍历
func isPalindrome(s string) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	s = strings.ToLower(s) // format

	left, right := 0, size-1
	for left < right {
		if !check(s[left]) {
			left++
			continue
		}
		if !check(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func check(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
		return true
	}
	return false
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
