/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	var res [][]int
	n, min := len(arr), math.MaxInt64
	sort.Ints(arr)
	for i := 0; i < n-1; i++ {
		diff := arr[i+1]-arr[i]
		if diff == min {
			res = append(res, []int{arr[i], arr[i+1]})
		} else if diff < min {
			min = diff
			res = [][]int{{arr[i], arr[i+1]}}
		}
	}
	return res
}
// @lc code=end

