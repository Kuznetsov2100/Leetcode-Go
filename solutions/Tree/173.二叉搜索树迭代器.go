/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.leftmostInorder(root) // 初始化栈，用于模拟中序遍历
	return iterator
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := len(this.stack)
	popnode := this.stack[n-1] // 栈顶元素始终是Next()函数需要的返回元素
	this.stack = this.stack[:n-1] // 模拟递归进入下一步
	if popnode.Right != nil { // 如果该节点的右子树不为空，调用辅助函数处理该节点的右子树
		this.leftmostInorder(popnode.Right)
	}
	return popnode.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 // 检查栈是否为空即可实现HasNext()
}

// 辅助函数： 将给定节点及节点的所有左子节点添加到栈中
func (this *BSTIterator) leftmostInorder(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

