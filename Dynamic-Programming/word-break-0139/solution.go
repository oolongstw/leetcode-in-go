package main

import "fmt"

// dp[i] 表示字符串前i个字符组成的字符串s[0..i−1] 是否能被空格拆分成若干个字典中出现的单词。
// dp[i] = dp[j] && checkStr(s[j,..,i]) 分为两部分 s[0..j-1] 和 s[j, i-1] 的正确性
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, val := range wordDict {
		wordDictSet[val] = true
	}
	dp := make([]bool, len(s)+1)
	// base case
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true // dp[i] 已经判断过，则退出该循环
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	dict := []string{"leet", "code"}
	fmt.Println(wordBreak("leetcode", dict))
}
