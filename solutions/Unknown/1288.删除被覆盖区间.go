/*
 * @lc app=leetcode.cn id=1288 lang=golang
 *
 * [1288] 删除被覆盖区间
 */

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	var deleteCount int
	length := len(intervals)
	// 数组按照每个区间的左端点升序排序,当左端点相等时，按右端点降序排序
	sort.Slice(intervals, func(i, j int) bool { 
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			if intervals[i][1] > intervals[j][1] {
				return true
			}
		}
		return false
	})
	
	for i := 0; i < length; {
		rightbound := intervals[i][1] // 初始化区间的右端点
		j := i+1
		for ; j < length && intervals[j][1] <= rightbound; j++ {
			deleteCount++
		}
		i = j
	}
	return length-deleteCount
}
// @lc code=end

