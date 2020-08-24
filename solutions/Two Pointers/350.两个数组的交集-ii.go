/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	m, n := len(nums1), len(nums2)
	index1, index2 := 0, 0
	for index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums2[index2] < nums1[index1] {
			index2++
		} else {
			res = append(res, nums1[index1])
			index1++
			index2++
		}
	}
	return res
}
// @lc code=end

