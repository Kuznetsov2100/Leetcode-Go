/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reorderList(head *ListNode)  {
// 	head = helper(head)
// }

// func helper(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	last := head
// 	var prelast *ListNode
// 	for last.Next != nil {
// 		prelast = last
// 		last = last.Next
// 	}
// 	prelast.Next = nil
// 	last.Next = helper(head.Next)
// 	head.Next = last
// 	return head
// }

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var deque []*ListNode // 用slice模拟双端队列存储链表所有节点
	for cur := head; cur != nil; {
		deque = append(deque, cur)
		cur = cur.Next
	}
	flag := 1  // flag为0时，从左边弹出，flag为1时从右边弹出
	p := deque[0] // 初始化指针
	deque = deque[1:]
	for len(deque) > 0 {
		var pop *ListNode
		if flag == 1 {
			 pop = deque[len(deque)-1]
			deque = deque[:len(deque)-1]
			flag = 0
		} else {
			pop = deque[0]
			deque = deque[1:]
			flag = 1
		}
		p.Next = pop
		p = p.Next
	}
	p.Next = nil // 最后一个节点指向nil
}
// @lc code=end

