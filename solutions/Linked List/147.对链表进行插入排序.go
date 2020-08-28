/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{ Next: head }
	tail, cur := head, head.Next
	for cur != nil {
		if cur.Val >= tail.Val {
			cur = cur.Next
			tail = tail.Next
		} else {
			pos := dummy
			for pos.Next.Val < cur.Val {
				pos = pos.Next
			}
			tail.Next = cur.Next
			cur.Next = pos.Next
			pos.Next = cur
			cur = tail.Next
		}
	}
	return dummy.Next
}
// @lc code=end

