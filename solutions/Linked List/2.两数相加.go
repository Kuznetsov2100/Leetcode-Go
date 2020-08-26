/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1,l2,0)
}

func helper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	x, y := 0, 0
	if l1 != nil {
		x = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		y = l2.Val
		l2 = l2.Next
	}
	sum := x + y + carry
	return &ListNode{ Val: sum%10, Next: helper(l1, l2, sum/10) }
}
// @lc code=end

