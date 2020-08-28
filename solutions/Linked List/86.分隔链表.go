/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	p1 := dummy1 // 比x小的节点组成链表p1
	p2 := dummy2 // 大于等于x的节点组成链表p2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p1.Next = dummy2.Next // p1尾部与p2的头部拼接
	p2.Next = nil // p2尾部终止
	return dummy1.Next // 返回p1头节点
}
// @lc code=end

