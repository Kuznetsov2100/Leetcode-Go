/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	var res []string
	if len(words) == 0 {
		return res
	}
	// 将不同行的字母分别映射成不同的数字
	keyboard := map[byte]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}
	for i := range words {
		cnt, j := 0, 0
		for ; j < len(words[i]); j++ {
			w := byte(words[i][j] | ' ') // 大写字母转成小写字母
			if j == 0 {
				cnt = keyboard[w] 
			} else {
				if keyboard[w] != cnt { // 如果该字符串后续字母的映射值不同于第一个，则break
					break
				}
			}
		}
		if j == len(words[i]) { // 如果j等于words[i]的长度，说明该字符串所有字母均处于同一行
			res = append(res, words[i]) // 将该字符串添加到结果数组
		}
	}
	return res
}
// @lc code=end

