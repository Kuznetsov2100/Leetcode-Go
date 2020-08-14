/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 将26个小写字母映射成26个质数，计算每个字符串的质数乘积，如果两个字符串的乘积相同，说明它们是异位词。
	var res [][]string
	alphaprime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,101}
	m := make(map[int][]string)
	for i := range strs {
		product := 1
		for j := range strs[i] {
			product *= alphaprime[strs[i][j]-97]
		}
		if _, ok := m[product]; ok {
			m[product] = append(m[product],strs[i])
		} else {
			m[product] = []string{strs[i]}
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
// @lc code=end

