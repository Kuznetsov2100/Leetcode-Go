/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strings"
import "strconv"
type Codec struct {
	null string
	sep string
	nodes []string
}

func Constructor() Codec {
	return Codec{null:"#", sep:","} 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
		var sb strings.Builder
		this.serializeHelper(root, &sb)
		return sb.String()
}

func (this *Codec) serializeHelper(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(this.null)
		sb.WriteString(this.sep)
		return 
	}
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(this.sep)
	this.serializeHelper(root.Left, sb)
	this.serializeHelper(root.Right, sb)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	this.nodes = strings.Split(data, this.sep)
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if len(this.nodes) == 0 {
		return nil
	}
	first := this.nodes[0]
	this.nodes = this.nodes[1:]
	if first == this.null {
		return nil
	}
	intVal, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}
	root := &TreeNode{Val:intVal}
	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

