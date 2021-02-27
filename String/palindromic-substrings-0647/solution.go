package main

import "fmt"

// 解题思路: 动态规划同 0005.最长回文子串
func countSubstrings(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	f := make([][]bool, n)
	// base case f[i][i] = true
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
		f[i][i] = true
	}
	// 状态转移方程 f[i][j] = f[i+1][j-1] && (s[i] == s[j]) f[i][j] 表示 s[i...j] 是否是回文串
	// 注意边界情况 子串长度小于2时直接赋值为true
	// 注意填表顺序 f[i][j] 依赖 f[i+1][j-1] 先计算i
	count := 0
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if s[i] == s[j] {
				if j-i < 3 {
					f[i][j] = true
				} else {
					f[i][j] = f[i+1][j-1]
				}
			}
			if f[i][j] {
				count++
			}
		}
	}

	return count
}

// 解题思路: 从一个点扩展出去
func countSubstrings2(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		helper(s, i, i, &res)   // aba 情况
		helper(s, i, i+1, &res) // aa 情况
	}
	return res
}

// left-- right++ 子串扩展
func helper(s string, left, right int, count *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*count++
		left--
		right++
	}
}

func main() {
	s := "babad"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstrings2(s))
}
