 /*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	var stackUp, stackDown []int // stackUp为单调递增栈，stackDown为单调递减栈
	n := len(nums)
	left, right := n-1, 0
	for i, j := 0, n-1; i < n && j >= 0; i,j = i+1, j-1 {
		// 从左到右找到最后一个破坏递增关系的元素的索引
		for len(stackUp) > 0 && nums[stackUp[len(stackUp)-1]] > nums[i] {
			if index := stackUp[len(stackUp)-1]; index < left {
				left = index
			}
			stackUp = stackUp[:len(stackUp)-1]
		}
		stackUp = append(stackUp, i)

		// 从右到左找到最后一个破坏递减关系的元素的索引
		for len(stackDown) > 0 && nums[stackDown[len(stackDown)-1]] < nums[j] {
			if index := stackDown[len(stackDown)-1]; index > right {
				right = index
			}
			stackDown = stackDown[:len(stackDown)-1]
		}
		stackDown = append(stackDown, j)
	}
	if right > left {
		return right-left+1
	}
	return 0
}
// @lc code=end

