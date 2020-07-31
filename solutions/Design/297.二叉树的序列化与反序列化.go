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
}

func Constructor() Codec {
	return Codec{} 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var serializeHelper func(root *TreeNode)
	serializeHelper = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("#")
			sb.WriteString(",")
			return 
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString(",")
		serializeHelper(root.Left)
		serializeHelper(root.Right)
	}
	serializeHelper(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	nodes := strings.Split(data, ",")
	var deserializeHelper func() *TreeNode
	deserializeHelper = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		first := nodes[0]
		nodes = nodes[1:]
		if first == "#" {
			return nil
		}
		intVal, _ := strconv.Atoi(first)
		return &TreeNode{
			Val: intVal,
			Left: deserializeHelper(),
			Right: deserializeHelper(),
		}
	}
	return deserializeHelper()
}
/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

