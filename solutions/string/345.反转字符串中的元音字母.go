/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
import "strings"
func reverseVowels(s string) string {
	str := []byte(s)
	for i,j := 0, len(str)-1; i < j;i,j = i+1,j-1 {
		for ;i < j && !strings.Contains("aAeEiIoOuU",string(str[i])); i++ {

		}
		for ;i < j && !strings.Contains("aAeEiIoOuU",string(str[j])); j-- {

		}
		if i < j {
			str[i], str[j] = str[j],str[i]
		}
	
	}
	return string(str)
}
// @lc code=end

