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
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
// @lc code=end

