/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
import "bytes"
import "strings"
func replaceWords(dict []string, sentence string) string {
	words := strings.Split(sentence, " ")
	wordtrie := Constructor()
	for i := range words {
		wordtrie.Insert(words[i])
	}
	root_word := make(map[string][]string)
	word_root := make(map[string][]string)
	for i := range dict {
		root_word[dict[i]] = wordtrie.StartsWith(dict[i])
	}
	for k, v := range root_word {
		for _, w := range v {
			word_root[w] = append(word_root[w], k)
		}
	}
	for i := range words {
		if length := len(word_root[words[i]]); length != 0 {
			words[i] = findmin(word_root[words[i]], length)
		}
	}
	return strings.Join(words, " ")
}

func findmin(strs []string, length int) string {
	res := strs[0]
	for i := 1; i < length; i++ {
		if len(strs[i]) < len(res) {
			res = strs[i]
		}
	}
	return res
}

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
func (this *Trie) StartsWith(prefix string) (results []string) {
	var b bytes.Buffer
	b.WriteString(prefix)
	this.collectPrefix(this.get(this.root, prefix, len(prefix), 0), &b, &results)
	return results
}

func (this *Trie) collectPrefix(x *node, prefix *bytes.Buffer, results *[]string ) {
	if x == nil {
		return
	}	
	if x.isString {
		*results = append(*results, prefix.String())
	}
	for c := 97; c <= 122; c++ {
		prefix.WriteByte(byte(c))
		this.collectPrefix(x.next[c-97], prefix, results)
		prefix.Truncate(prefix.Len()-1)
	}
}
// @lc code=end

