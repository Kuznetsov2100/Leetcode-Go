/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	existed := make(map[int]bool)
	for i := range nums1 {
		existed[nums1[i]] = true
	}
	for i := range nums2 {
		if existed[nums2[i]] {
			res = append(res, nums2[i])
			existed[nums2[i]] = false
		}
	}
	return res
}
// @lc code=end

