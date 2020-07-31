/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	guard := &ListNode{Next:head} // 哨兵节点
	slow, fast := guard, guard // 快慢指针指向哨兵节点(头节点之前的虚拟节点)
	for n >= 0 { // 快指针先走n+1步
		fast = fast.Next
		n--
	}
	for fast != nil { // 快慢指针同步前进，当快指针指向nil时，慢指针指向倒数第n+1个数
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next // 慢指针指向倒数第n-1个数，就删除了倒数第n个数
	return guard.Next
}
// @lc code=end

