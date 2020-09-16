/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	// 排序+双指针
	// 升序排序以后，假如a<=b<=c ,当且仅当a+b > c，a,b，c三条边能组成三角形
	var res int
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for k := n-1; k >= 2; k-- { // 固定最长的边nums[k]
		i := 0 // nums[i]为最短边
		j := k-1 // 以nums[k-1]作为剩下的一条边
		for i < j {
			if nums[i] + nums[j] > nums[k] {
				res += j-i // 说明从nums[i]...nums[j-1]都满足条件，构成三角形的三元组个数为j-i
				j-- // j左移一位，继续下一轮判断
			} else {
				i++ // i右移一位，继续下一轮判断
			}
		}
	}
	return res
}
// @lc code=end

