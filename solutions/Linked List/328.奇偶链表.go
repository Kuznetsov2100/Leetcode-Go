/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd, even := head, head.Next
	node := head.Next
	for odd != nil && even != nil {
		if odd.Next.Next != nil {
			odd.Next = odd.Next.Next
			odd = odd.Next
			even.Next = odd.Next
			even = even.Next
		} else {
			break
		}
	}
	odd.Next = node
	return head
}
// @lc code=end

