package main

import (
	"fmt"
	"sort"
)

// 解题思路 每个字符串中的字符出现的频率
func groupAnagrams(strs []string) [][]string {
	hash := map[[26]int][]string{}
	for _, s := range strs {
		tmp := [26]int{}
		for i := 0; i < len(s); i++ {
			tmp[s[i]-'a']++
		}
		hash[tmp] = append(hash[tmp], s)
	}
	res := [][]string{}
	for _, val := range hash {
		res = append(res, val)
	}
	return res
}

type bytes []byte

// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// 方法二 解题思路: 字符串排序
func groupAnagrams2(strs []string) [][]string {
	ret := [][]string{}
	m := make(map[string][]string)
	for _, str := range strs {
		kByte := bytes(str)
		sort.Sort(kByte) // 将字符串排序
		k := string(kByte)
		m[k] = append(m[k], str)
	}
	for _, val := range m {
		ret = append(ret, val)
	}
	return ret
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
}
