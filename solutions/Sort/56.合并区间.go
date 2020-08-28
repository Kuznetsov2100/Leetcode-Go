/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 数组按照每个区间的左端点升序排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var res [][]int
	for i := 0; i < len(intervals); {
		rightbound := intervals[i][1] // 初始化区间的右端点
		j := i+1
		for j < len(intervals) && intervals[j][0] <= rightbound {
			if intervals[j][1] > rightbound {  // 拓展区间的右端点
				rightbound = intervals[j][1]
			}
			j++
		}
		res = append(res, []int{intervals[i][0], rightbound}) // 将一个合并后的区间加入到结果数组
		i = j
	}
	return res
}
// @lc code=end

