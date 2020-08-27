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
func reverseBetween(head *ListNode, m int, n int) *ListNode  {
	dummy := &ListNode{ Next:head } // 定义一个dummy节点: 考虑到m=1的情况
	prem := dummy
	for i := 1; i <= m-1; i++ {
		prem = prem.Next
	}
	cur := prem.Next // cur处于位置m
	var pre *ListNode
	for i := m; i <= n; i++ { // 位置m到位置n之间的节点反转
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	prem.Next.Next = cur // 位置m需要指向位置n之后的节点
	prem.Next = pre // prem需要指向区间反转后的头节点
	return dummy.Next // 返回头节点
}

// 个人解法
// func reverseBetween(head *ListNode, m int, n int) *ListNode {
// 	if head == nil || m == n {
// 		return head
// 	}
// 	prem, posn := head, head
// 	for ;m > 2; m-- {
// 		prem = prem.Next
// 	}
// 	for ;n > 1; n-- {
// 		posn = posn.Next
// 	}
// 	postn := posn.Next
// 	posn.Next = nil

// 	var newhead *ListNode
// 	if m == n  {
// 		newhead = reverseList(prem)
// 	} else {
// 		newhead = head
// 		prem.Next = reverseList(prem.Next)	
// 	}
// 	cur := newhead
// 	for cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = postn
// 	return newhead
// }

// func reverseList(head *ListNode) *ListNode {
// 	// pre记录前驱，next记录后继，不断更新cur.Next = pre
// 	var pre,next *ListNode
// 	cur := head
// 	for cur != nil {
// 		next = cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}
// 	return pre
// }

// @lc code=end

