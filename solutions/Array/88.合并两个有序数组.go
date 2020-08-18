/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1 := m-1 // 指向nums1最后一个位置
	p2 := n-1 // 指向nums2最后一个位置
	p3 := m+n-1 // 指向结果数组的最后一个位置
	for p2 >= 0 { 
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
}
// @lc code=end

