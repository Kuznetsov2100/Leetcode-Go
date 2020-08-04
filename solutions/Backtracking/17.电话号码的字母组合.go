/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[byte]string {
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var res []string
	var backtrack func(path string, index int)
	backtrack = func(path string, index int) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		for _, s := range m[digits[index]] {
			backtrack(path+string(s), index+1)
		}
	}
	backtrack("",0)
	return res
}
// @lc code=end

