package main

import "fmt"

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if prefix == "" {
			break
		}
		pos := helper(prefix, strs[i])
		prefix = prefix[:pos]
	}
	return prefix
}

func helper(a, b string) int {
	l := min(len(a), len(b))
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return l
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	strs := []string{"aab", "abc", "abcc"}
	fmt.Println(longestCommonPrefix(strs))
}
