package main

import "fmt"

// 解题思路: 最长回文子串
func longestPalindrome(s string) string {
	size := len(s)
	res := ""
	// 从最长的长度开始寻找
	for length := size; length > 0; length-- {
		for i := 0; i < size-length+1; i++ {
			sub := s[i : i+length]
			if isPalindrome(sub, length) {
				return sub
			}
		}
	}
	return res
}

// 判断字符是否是回文
func isPalindrome(s string, size int) bool {
	left, right := 0, size-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "babab"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome2(s))
}
