/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 import "math/rand"
type Solution struct {
        head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
        return Solution{head:head}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
        // 蓄水池抽样算法
        var res int
        i := 1
        cur := this.head        
        for cur != nil {
                if rand.Intn(i) == 0 {
                    res = cur.Val
                }
                cur = cur.Next
                i++
        }
        return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

