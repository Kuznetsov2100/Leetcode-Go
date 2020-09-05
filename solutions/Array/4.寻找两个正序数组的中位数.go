/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := merge(nums1, nums2)
	n := len(arr)
	mid := (n-1)/2
	if n % 2 == 1 {
		return float64(arr[mid])
	}
	return float64(arr[mid]+arr[mid+1])*0.5
}

func merge(nums1, nums2 []int) []int {
	index1, index2 := 0, 0
	m, n := len(nums1), len(nums2)
	var res []int
	for index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			res = append(res, nums1[index1])
			index1++
		} else {
			res = append(res, nums2[index2])
			index2++
		}
	}
	for index1 < m {
		res = append(res, nums1[index1])
		index1++
	}
	for index2 < n {
		res = append(res, nums2[index2])
		index2++
	}
	return res
}
// @lc code=end

