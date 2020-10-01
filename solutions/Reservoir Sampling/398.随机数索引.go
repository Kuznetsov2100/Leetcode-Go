/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
import "math/rand"
type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{nums:nums}
}


func (this *Solution) Pick(target int) int {
	// 蓄水池抽样算法
	var res int
	j := 1
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if rand.Intn(j) == 0 {
				res = i
			}
			j++
		}
	}
	return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

