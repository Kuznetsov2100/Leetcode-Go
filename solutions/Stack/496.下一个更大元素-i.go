/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	var stack []int // 单调递减栈
	m := make(map[int]int)
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			element := stack[len(stack)-1] // num为右边第一个比栈顶元素大的元素
			stack = stack[:len(stack)-1] // 出栈
			m[element] = num // 建立映射关系
		}
		stack = append(stack, num) // 入栈
	}

	// 如果stack非空，因为是单调递减栈，
	// 栈内元素从栈底到栈顶按递减关系排列
	// 栈内没有元素能找到右边第一个比自己大的元素
	for _, num := range stack {
		m[num] = -1
	}
	for i, num := range nums1 {
		res[i] = m[num]
	}
	return res
}

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	len1, len2 := len(nums1), len(nums2)
// 	m := make(map[int]int)
// 	for i := 0; i < len2; i++ {
// 		m[nums2[i]] = -1
// 		for j := i+1; j < len2; j++ {
// 			if nums2[j] > nums2[i] {
// 				m[nums2[i]] = nums2[j]
// 				break
// 			}
// 		}
// 	}
// 	res := make([]int, len1)
// 	for i := 0; i < len1; i++ {
// 		res[i] = m[nums1[i]]
// 	}
// 	return res
// }
// @lc code=end

