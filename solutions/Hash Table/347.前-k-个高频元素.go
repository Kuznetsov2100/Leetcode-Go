/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
import "sort"
func topKFrequent(nums []int, k int) []int {
	type foo struct {
		number int
		frequency int
	}
	var arr []foo
	var res []int
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for k, v := range m {
		arr = append(arr, foo{number: k, frequency: v})
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].frequency > arr[j].frequency })
	for i := 0; i < k; i++ {
		res = append(res,arr[i].number)
	}
	return res
}
// @lc code=end

