/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var b bytes.Buffer
	for head != nil {
		if head.Val == 0 {
			b.WriteByte('0')
		} else {
			b.WriteByte('1')
		}
		head = head.Next
	}
	res, _ := strconv.ParseInt(b.String(), 2, 32)
	return int(res)
}
// @lc code=end

