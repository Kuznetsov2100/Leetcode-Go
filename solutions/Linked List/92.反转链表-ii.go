/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	prem, posn := head, head
	for ;m > 2; m-- {
		prem = prem.Next
	}
	for ;n > 1; n-- {
		posn = posn.Next
	}
	postn := posn.Next
	posn.Next = nil

	var newhead *ListNode
	if m == n  {
		newhead = reverseList(prem)
	} else {
		newhead = head
		prem.Next = reverseList(prem.Next)	
	}
	cur := newhead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = postn
	return newhead
}

func reverseList(head *ListNode) *ListNode {
	// pre记录前驱，next记录后继，不断更新cur.Next = pre
	var pre,next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
// @lc code=end

