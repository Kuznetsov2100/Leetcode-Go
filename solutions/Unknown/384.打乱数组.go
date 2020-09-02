/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
import "math/rand"
type Solution struct {
	original []int
}


func Constructor(nums []int) Solution {
	return Solution{original:nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.original)
	nums := make([]int, n)
	copy(nums, this.original)
	rand.Shuffle(n, func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	return nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

