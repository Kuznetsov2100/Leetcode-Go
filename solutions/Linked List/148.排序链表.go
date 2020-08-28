/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var premid *ListNode // 中点之前的节点
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		premid = slow
		slow = slow.Next
	}
	mid := premid.Next
	premid.Next = nil // 断开
	return merge(sortList(head), sortList(mid)) // 左半边链表head到premid,右半边mid到最后
}

// 迭代版本
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

// 递归版本
// func merge(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		l1.Next = merge(l1.Next, l2)
// 		return l1
// 	} else {
// 		l2.Next = merge(l1, l2.Next)
// 		return l2
// 	}
// }
// @lc code=end

