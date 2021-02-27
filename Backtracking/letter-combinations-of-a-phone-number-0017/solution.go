package main

import (
	"fmt"
)

var phone map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var (
	ret  []string
	tmpl string
)

// 电话号码的字母组合
func letterCombinations(digits string) []string {
	ret = []string{}
	if digits == "" {
		return ret
	}
	tmpl = digits
	dfs(0, "")
	return ret
}

func dfs(start int, prefix string) {
	if start == len(tmpl) {
		ret = append(ret, prefix)
		return
	}
	digit := string(tmpl[start])
	letters := phone[digit]
	// 下一次 dfs 起始的 index
	for i := 0; i < len(letters); i++ {
		dfs(start+1, prefix+string(letters[i]))
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
