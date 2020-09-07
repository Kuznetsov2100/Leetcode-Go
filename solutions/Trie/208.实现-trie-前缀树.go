/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isEnd bool
	children [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, ch := range word {
		index := ch-'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{} 
		}
		node = node.children[index]
	}
	node.isEnd =  true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

