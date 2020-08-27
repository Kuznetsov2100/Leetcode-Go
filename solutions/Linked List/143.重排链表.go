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
 // 第四版: 无需额外空间
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    // 快慢指针寻找链表中点
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    p := slow.Next // slow为中点， p为中点后的节点 
    slow.Next = nil // 切断中点后的节点
    
    // 反转以p为头节点的链表
    var pre *ListNode	
	for cur := p; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
    }

    // 将反转后的链表与前半段链表交替链接
    p1 := head
    p2 := pre
    for p2 != nil {
        p1next := p1.Next
        p2next := p2.Next
        p1.Next = p2
        p2.Next = p1next
        p1 = p1next
        p2 = p2next
    }
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

// 第三版: 依然是用slice储存所有链表节点，但不需要弹出操作，
// 使用左右指针来完成重排操作
// func reorderList(head *ListNode) {
//     if head == nil || head.Next == nil {
//         return
//     }
//     var arr []*ListNode
//     for cur := head; cur != nil; {
//         arr = append(arr, cur)
//         cur = cur.Next
//     }
//     left, right := 0, len(arr)-1
//     for left < right  {
//         arr[left].Next = arr[right]
//         left++
//         arr[right].Next = arr[left]
//         right--
//     }
//     arr[left].Next = nil
// }
// @lc code=end

