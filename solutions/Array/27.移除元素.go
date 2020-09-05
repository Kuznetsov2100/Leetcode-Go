/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	// 定义nums[0,index)为没有元素等于val的数组区间
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index // index-1-0+1
}
// func removeElement(nums []int, val int) int {
// 	slow := 0
// 	for fast := 0; fast < len(nums); fast++ {
// 		if nums[fast] != val {
// 			nums[slow] = nums[fast]
// 			slow++
// 		}
// 	}
// 	return slow
// }
// @lc code=end

