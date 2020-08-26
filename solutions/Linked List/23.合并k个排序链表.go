/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// 分治合并
// k个链表两两配对，进行第一轮合并，结束后k个链表被合并成k/2个链表
// k/2个链表依然两两配对，进行第二轮合并，结束后k/2个链表被合并成k/4个链表
// 重复上述过程，进行log(k)次合并，完成总体合并工作
func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right - left)/2
	return mergeTwoLists(merge(lists,left, mid), merge(lists, mid+1, right))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
// @lc code=end

