/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	m := make(map[int]int)
	for i := 0; i < len2; i++ {
		m[nums2[i]] = -1
		for j := i+1; j < len2; j++ {
			if nums2[j] > nums2[i] {
				m[nums2[i]] = nums2[j]
				break
			}
		}
	}
	res := make([]int, len1)
	for i := 0; i < len1; i++ {
		res[i] = m[nums1[i]]
	}
	return res
}
// @lc code=end

