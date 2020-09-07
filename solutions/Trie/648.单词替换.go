/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dict []string, sentence string) string {
	wordtrie := Constructor()
	for _, word := range dict {
		wordtrie.Insert(word)
	}
	
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if index := wordtrie.findShortestPrefix(w); index != -1 {
			words[i] = w[:index]
		}
	}
	return strings.Join(words, " ")
}

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

func (this *Trie) findShortestPrefix(word string) int {
	node := this
	res := 0
	for _, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return -1
		}
		res++
		if node.isEnd {
			return res
		}
	}
	return res
}
// @lc code=end

