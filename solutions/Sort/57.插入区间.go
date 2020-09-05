/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
// 思路：原区间按照起始端点排序的，具有有序性
// 首先利用二分搜索找到新区间应该插入到intervals的位置，
// 然后用leetcode-56题的merge函数处理即可
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	pos := searchInsert(intervals, newInterval[0])
	return merge(append(intervals[:pos], append([][]int{newInterval}, intervals[pos:]...)...))
}

func merge(intervals [][]int) [][]int {
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

func searchInsert(nums [][]int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if value := nums[mid][0]; value < target {
			left = mid + 1
		} else if value > target {
			right = mid - 1
		} else if value == target {
			return mid
		}
	}
	return left
}
// @lc code=end

