package main

import "fmt"

// 解题思路: hash 判断两个字符串中的字符数量是否相等
func isAnagram(s string, t string) bool {
	hash := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		hash[t[i]]--
	}
	for _, val := range hash {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "car", "car"
	fmt.Println(isAnagram(s, t))
}
