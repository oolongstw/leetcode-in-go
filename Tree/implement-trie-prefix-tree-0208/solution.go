package main

const N = 26 // 小写字母 最多26个

// 实现 Trie (前缀树)
type Trie struct {
	isEnd bool
	next  [N]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if !this.containsKey(word[i]) {
			this.next[word[i]-'a'] = &Trie{} // 新节点
		}
		this = this.next[word[i]-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if !this.containsKey(word[i]) {
			return false
		}
		this = this.next[word[i]-'a']
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if !this.containsKey(prefix[i]) {
			return false
		}
		this = this.next[prefix[i]-'a']
	}
	return true
}

func (this *Trie) containsKey(key byte) bool {
	return this.next[key-'a'] != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
