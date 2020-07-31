/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast { // slow 走k步，fast走2k步，slow和fast在环的点m处相遇，点m距离环的起点距离为m，环的长度为k
			break
		}
	}
	if fast == nil || fast.Next == nil { 
		return nil // 链表无环
	}
	slow = head
	for slow != fast { // slow和fast同步走k-m步后在环的起点相遇
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
// @lc code=end

