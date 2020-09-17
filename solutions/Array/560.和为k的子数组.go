/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	var res, sum int
	m := make(map[int]int) // key:前缀和的值 value:出现次数
	m[0] = 1 // 开始遍历前， 记录前缀和为0的次数为1
	for i := range nums {
		sum += nums[i] // sum为当前的前缀和
		if count, ok := m[sum-k]; ok { // 如果在map中找到了某个前缀和x,sum-x==k，表明找到了count个连续子数组的和为k
			res += count
		}
		m[sum]++ // 该前缀和出现的次数+1
	}
	return res
}
// @lc code=end

