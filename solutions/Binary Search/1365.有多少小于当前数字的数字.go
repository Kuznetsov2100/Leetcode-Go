/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)
	for i, num := range nums {
		res[i] = sort.SearchInts(arr, num)
	}
	return res
}
// @lc code=end

