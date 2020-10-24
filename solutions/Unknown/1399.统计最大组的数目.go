/*
 * @lc app=leetcode.cn id=1399 lang=golang
 *
 * [1399] 统计最大组的数目
 */

// @lc code=start
func countLargestGroup(n int) int {
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum, num := 0, i
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		m[sum]++
	}
	max, res := 0, 0
	for _, count := range m {
		if count > max {
			max = count
			res = 1
		} else if count == max {
			res++
		}
	}
	return res
}
// @lc code=end

