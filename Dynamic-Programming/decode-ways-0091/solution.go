package main

import "fmt"

// f[i] = f[i-2] + f[i-1] 注意 0 的边界情况
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	size := len(s)
	f := make([]int, size+1)
	f[0], f[1] = 1, 1 // base case
	for i := 2; i <= size; i++ {
		if s[i-1] == '0' {
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			} else {
				f[i] = f[i-2] // 前一位为 1 or 2的情况
				continue
			}
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			// 处理 1x 和 20~26 的数字
			f[i] = f[i-1] + f[i-2]
		} else {
			f[i] = f[i-1]
		}
	}
	return f[size]
}

func main() {
	sample := "102"
	fmt.Println(numDecodings(sample))
}
