/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	length := 0
	var oldLast *ListNode
	for cur := head; cur != nil; {
		oldLast = cur
		cur = cur.Next
		length++
	}
	k = k % length
	if k == 0 {
		return head
	}
	newhead := head
	var newLast *ListNode
	for i := 0; i < length-k; i++ {
		newLast = newhead
		newhead = newhead.Next	
	}
	newLast.Next = nil
	oldLast.Next = head
	return newhead	
}
// @lc code=end

