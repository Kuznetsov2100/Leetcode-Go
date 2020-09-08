/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start
type MagicDictionary struct {
	isEnd bool
	children [26]*MagicDictionary
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	for _, word := range dictionary {
		this.addWord(word)
	}
}	

func (this *MagicDictionary) addWord(word string)  {
	node := this
	for _, ch := range word {
		index := ch-'a'
		if node.children[index] == nil {
			node.children[index] = &MagicDictionary{} 
		}
		node = node.children[index]
	}
	node.isEnd =  true
}



func (this *MagicDictionary) Search(searchWord string) bool {
	return this.dfs(this, searchWord, len(searchWord), 0, 1)
}

func (this *MagicDictionary) dfs(node *MagicDictionary, searchWord string, wordlen, i, quota int) bool {
	if quota < 0 {
		return false
	}	
	if i == wordlen {
		return quota == 0 && node.isEnd
	}
	
	index := int(searchWord[i]-'a')
	for j := 0; j < 26; j++ {
		if node.children[j] != nil {
			if j == index {
				if this.dfs(node.children[j], searchWord, wordlen, i+1, quota) {
					return true
				}
			} else {
				if this.dfs(node.children[j],searchWord, wordlen, i+1, quota-1) {
					return true
				}
			}
		}
	}
	return false
}



/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

