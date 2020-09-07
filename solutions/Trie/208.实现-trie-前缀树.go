/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
import "bytes"
type Trie struct {
	root *node
}

type node struct {
	isString bool
	next [26]*node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.root = this.insert(this.root, word, len(word), 0)
}

func (this *Trie) insert(x *node, key string, keylength, d int) *node {
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


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	x := this.get(this.root, word, len(word), 0)
	if x == nil {
		return false
	}
	return x.isString
}

func (this *Trie) get(x *node, key string, keylength, d int) *node {
	if x == nil {
		return nil
	}
	if d == keylength {
		return x
	}
	return this.get(x.next[key[d]-97], key, keylength, d+1)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var b bytes.Buffer
	b.WriteString(prefix)
	return this.collectPrefix(this.get(this.root, prefix, len(prefix), 0), &b)
}

func (this *Trie) collectPrefix(x *node, prefix *bytes.Buffer) bool {
	if x == nil {
		return false
	}	
	if x.isString {
		return true
	}
	for c := 97; c <= 122; c++ {
		prefix.WriteByte(byte(c))
		if this.collectPrefix(x.next[c-97], prefix) {
			return true
		}
		prefix.Truncate(prefix.Len()-1)
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

