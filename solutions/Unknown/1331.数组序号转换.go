/*
 * @lc app=leetcode.cn id=1331 lang=golang
 *
 * [1331] 数组序号转换
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	nums := make([]int, n)
	copy(nums, arr)
	sort.Ints(nums)
	m := make(map[int]int)
	rank := 1
	for i := range nums {
		if i == 0 {
			m[nums[i]] = rank
			continue
		}
		if nums[i] > nums[i-1] {
			rank++
			m[nums[i]] = rank
		} else if nums[i] == nums[i-1] {
			m[nums[i]] = rank
		}
	}
	res := make([]int, n)
	for i := range arr {
		res[i] = m[arr[i]]
	}
	return res
}
// @lc code=end

