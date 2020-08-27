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
 // 第三版: 依然是用slice储存所有链表节点，但不需要弹出操作，
 // 使用左右指针来完成重排操作
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    var arr []*ListNode
    for cur := head; cur != nil; {
        arr = append(arr, cur)
        cur = cur.Next
    }
    left, right := 0, len(arr)-1
    for left < right  {
        arr[left].Next = arr[right]
        left++
        arr[right].Next = arr[left]
        right--
    }
    arr[left].Next = nil
}

// 第一版：递归解法 效率十分低下
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

// 第二版:双端队列, 时间上进步很大，需要一定的空间来储存链表所有节点
// func reorderList(head *ListNode) {
// 	if head == nil || head.Next == nil {
// 		return
// 	}
// 	var deque []*ListNode // 用slice模拟双端队列存储链表所有节点
// 	for cur := head; cur != nil; {
// 		deque = append(deque, cur)
// 		cur = cur.Next
// 	}
// 	flag := 1  // flag为0时，从左边弹出，flag为1时从右边弹出
// 	p := deque[0] // 初始化指针
// 	deque = deque[1:]
// 	for len(deque) > 0 {
// 		var pop *ListNode
// 		if flag == 1 {
// 			 pop = deque[len(deque)-1]
// 			deque = deque[:len(deque)-1]
// 			flag = 0
// 		} else {
// 			pop = deque[0]
// 			deque = deque[1:]
// 			flag = 1
// 		}
// 		p.Next = pop
// 		p = p.Next
// 	}
// 	p.Next = nil // 最后一个节点指向nil
// }
// @lc code=end

