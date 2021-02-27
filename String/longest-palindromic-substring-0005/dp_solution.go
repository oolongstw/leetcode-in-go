package main

// 解题思路: 动态规划
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	f := make([][]bool, n)
	// base case f[i][i] = true
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
		f[i][i] = true
	}
	// 状态转移方程 f[i][j] = f[i+1][j-1] && (s[i] == s[j]) f[i][j] 表示 s[i...j] 是否是回文串
	// 注意边界情况 子串长度小于2时直接赋值为true
	// 注意填表顺序 f[i][j] = f[i+1][j-1] 先计算i
	maxLen, start := 1, 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				f[i][j] = false
			} else {
				if j-i < 3 {
					f[i][j] = true
				} else {
					f[i][j] = f[i+1][j-1]
				}
			}
			if f[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}
