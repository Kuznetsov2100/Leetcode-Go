/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
import "bytes"
type MapSum struct {
	isEnd bool
	val int
	children [26]*MapSum
}


/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}


func (this *MapSum) Insert(key string, val int)  {
	node := this
	for _, ch := range key {
		index := ch-'a'
		if node.children[index] == nil {
			node.children[index] = &MapSum{} 
		}
		node = node.children[index]
	}
	node.isEnd =  true
	node.val = val
}	

func (this *MapSum) get(x *MapSum, key string, keylength, d int) *MapSum {
	if x == nil {
		return nil
	}
	if d == keylength {
		return x
	}
	return this.get(x.children[key[d]-'a'], key, keylength, d+1)
}

func (this *MapSum) Sum(prefix string) int {
	var b bytes.Buffer
	var res int
	b.WriteString(prefix)
	this.collectPrefix(this.get(this, prefix, len(prefix), 0), &b, &res)
	return res
}

func (this *MapSum) collectPrefix(x *MapSum, prefix *bytes.Buffer, result *int) {
	if x == nil {
		return
	}
	if x.isEnd {
		*result += x.val
	}

	for c := 97; c <= 122; c++ {
		prefix.WriteByte(byte(c))
		this.collectPrefix(x.children[c-97], prefix, result)
		prefix.Truncate(prefix.Len() - 1)
	}
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

