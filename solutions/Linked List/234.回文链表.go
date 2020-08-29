/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head // 快慢指针找中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 反转以中点为头节点的后半段链表
	var pre *ListNode // 反转以后的链表的头节点
	for slow != nil {
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	// 依次比较head和pre值是否相同
	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}
// @lc code=end

