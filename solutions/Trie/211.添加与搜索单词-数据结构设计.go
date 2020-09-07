/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type WordDictionary struct {
	isEnd bool
	children [26]*WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	node := this
	for _, ch := range word {
		index := ch-'a'
		if node.children[index] == nil {
			node.children[index] = &WordDictionary{} 
		}
		node = node.children[index]
	}
	node.isEnd =  true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var b bytes.Buffer
	return this.collectMatch(this, &b, word, len(word))
}

func (this *WordDictionary) collectMatch(x *WordDictionary, prefix *bytes.Buffer, pattern string, plen int) bool {
	if x == nil {
		return false
	}	
	d := prefix.Len()
	if d == plen {
		if x.isEnd {
			return true
		}
		return false
	}

	if c := pattern[d]; c == '.' {
		for ch := 97; ch <= 122; ch++ {
			prefix.WriteByte(byte(ch))
			if this.collectMatch(x.children[ch-97], prefix, pattern, plen) {
				return true
			}
			prefix.Truncate(prefix.Len()-1)
		}
	} else {
		prefix.WriteByte(c)
		if this.collectMatch(x.children[c-97], prefix, pattern, plen) {
			return true
		}
		prefix.Truncate(prefix.Len()-1)
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

