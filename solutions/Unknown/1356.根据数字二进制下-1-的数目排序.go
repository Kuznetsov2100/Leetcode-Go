/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
import "sort"
func sortByBits(arr []int) []int {
	m := make(map[int]int)
	for i := range arr {
		m[arr[i]] = hammingWeight(arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		if m[arr[i]] < m[arr[j]] {
			return true
		} else if m[arr[i]] == m[arr[j]] {
			if arr[i] < arr[j] {
				return true
			}
		}
		return false
	})
	return arr
}

func hammingWeight(num int) int {
	var res int
	for num != 0 {
		num = num & (num-1) // n & (n-1)消除数字n的二进制表示中的最后一个1
		res++
	}
	return res
}
// @lc code=end

