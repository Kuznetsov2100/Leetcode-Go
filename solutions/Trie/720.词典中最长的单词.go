/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) string {
	var res string
	wordtrie := Constructor()
	for _, word := range words {
		wordtrie.Insert(word)
	}
	sort.Slice(words, func(i, j int) bool {
		m, n := len(words[i]), len(words[j])
		if m > n {
			return true
		} else if m == n {
			if  words[i] < words[j] {
				return true
			}
		}
		return false
	})
	for _, word := range words {
		i := 0
		length := len(word)
		for ; i < length; i++ {
			if !wordtrie.Search(word[:i+1]) {
				break
			}
		}
		if i == length {
			res = word
			break
		}
	}
	return res
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
// @lc code=end

