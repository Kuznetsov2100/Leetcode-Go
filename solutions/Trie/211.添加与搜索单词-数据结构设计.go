/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
import "bytes"
type WordDictionary struct {
	root *node

}

type node struct {
	isString bool
	next [26]*node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	this.root = this.insert(this.root, word, len(word), 0)
}

func (this *WordDictionary) insert(x *node, key string, keylength, d int) *node {
	if x == nil {
		x = &node{}
	}
	if d == keylength {
		x.isString = true
	} else {
		x.next[key[d]-97] = this.insert(x.next[key[d]-97], key, keylength, d+1)
	}
	return x
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var b bytes.Buffer
	return this.collectMatch(this.root, &b, word, len(word))
}

func (this *WordDictionary) collectMatch(x *node, prefix *bytes.Buffer, pattern string, plen int) bool {
	if x == nil {
		return false
	}	
	d := prefix.Len()
	if d == plen {
		if x.isString {
			return true
		}
		return false
	}

	if c := pattern[d]; c == '.' {
		for ch := 97; ch <= 122; ch++ {
			prefix.WriteByte(byte(ch))
			if this.collectMatch(x.next[ch-97], prefix, pattern, plen) {
				return true
			}
			prefix.Truncate(prefix.Len()-1)
		}
	} else {
		prefix.WriteByte(c)
		if this.collectMatch(x.next[c-97], prefix, pattern, plen) {
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

