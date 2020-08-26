/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // 先排序
	res := nums[0] + nums[1] + nums[2] // 初始值
	for  i := 0; i < len(nums); i++ { // 每次固定一个nums[i],剩下的两个数使用双指针来寻找
		left, right := i+1, len(nums)-1 // 左右指针
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(res-target) { // 如果sum更靠近target,更新res
				res = sum
			}
			if sum < target { // 使sum变大
				left++
			} else if sum > target { // 使sum变小
				right--
			} else {
				return res // sum等于target，直接返回结果
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

