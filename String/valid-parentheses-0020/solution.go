package main

import "fmt"

// 解题思路 利用栈数据结构
func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		if m[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1] // 栈顶元素匹配 -> 从栈中去除
			continue
		} else {
			stack = append(stack, s[i]) // 其他情况入栈
		}
	}
	return len(stack) == 0
}

func main() {
	sample := "{{}{}}"
	fmt.Println(isValid(sample))
}
