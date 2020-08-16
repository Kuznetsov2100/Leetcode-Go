/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	s1map, s2map := [26]int{}, [26]int{} // s1map记录s1各字符出现频率
	for i := 0; i < m; i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++ // 记录s2第一个窗口中各字符出现频率
	}
	for i := 0; i < n-m; i++ {
		if s1map == s2map { // golang中，数组可直接使用 == 进行比较
			return true
		}
		s2map[s2[i+m]-'a']++ // 添加一个新字符进入窗口
		s2map[s2[i]-'a']--  // 移除窗口最左侧字符
	}
	return s1map == s2map // 比较最后一个窗口和s1map是否相等
}
// @lc code=end

