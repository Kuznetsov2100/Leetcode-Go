/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	n := len(nums)
	var res []string
	for left, right := 0, 0; right < n; right++ {
		// 尝试拓展区间
		if right < n-1 && nums[right+1]-nums[right] == 1 {
			continue
		}

		// 添加相应区间
		if left == right { // 单个数字构成的区间
			res = append(res, strconv.Itoa(nums[left]))
		} else {
			res = append(res, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[right]))
		}
		left = right + 1 // 新的区间的左端索引
	}
	return res
}
// @lc code=end

