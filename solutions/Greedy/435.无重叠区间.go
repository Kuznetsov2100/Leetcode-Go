/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	var res int
	if len(intervals) == 1 {
		return res
	}
	// 数组按照每个区间的左端点升序排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 0; i < len(intervals); {
		rightbound := intervals[i][1] // 初始化区间的右端点
		j := i+1
		for j < len(intervals) && intervals[j][0] < rightbound {
			res++
			if intervals[j][1] < rightbound { // 最多的不重叠区间意味着合并区间的次数最少，所以rightbound尽量最小
				rightbound = intervals[j][1]
			}
			j++
		}
		i = j
	}
	return res
}
// @lc code=end

