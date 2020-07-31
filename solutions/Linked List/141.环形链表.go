/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 解释：快指针每次前进两步，慢指针每次前进一步。如果不含有环，快指针最终会遇到nil，说明链表不含环；
	// 如果含有环，快指针最终会超慢指针一圈，和慢指针相遇，说明链表含有环。
	fast, slow := head,head
	for fast != nil  && fast.Next != nil { // fast.Next != nil 是为了确保下方的fast.Next.Next是合法的指针
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
// @lc code=end

