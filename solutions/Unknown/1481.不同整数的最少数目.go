/*
 * @lc app=leetcode.cn id=1481 lang=golang
 *
 * [1481] 不同整数的最少数目
 */

// @lc code=start
func findLeastNumOfUniqueInts(arr []int, k int) int {
	// 统计数字出现频率---> 排序----> 贪心思想
	var res int
	var counts []int
	n := len(arr)

	if k == n {
		return 0
	}
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}	
	for _, v := range m {
		counts = append(counts, v)
	}
	sort.Slice(counts, func(i, j int) bool {return counts[i] > counts[j] })

	sum := 0
	for _, num := range counts {
		if sum >= n-k {
			break
		}
		sum += num
		res++
	}
	return res	
}
// @lc code=end

