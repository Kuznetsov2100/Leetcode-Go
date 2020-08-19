/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
import "strings"
func intToRoman(num int) string {
	// 把阿拉伯数字与罗马数字可能出现的所有情况和对应关系
    // 按照阿拉伯数字的大小降序排列，根据贪心算法的思想，每次选择尽可能大的符号来表示
	arr := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	romans := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	var res strings.Builder
	for i := 0; i < len(arr); i++ {
		for num >= arr[i] {
			res.WriteString(romans[i])
			num -= arr[i]
		}
	}
	return res.String()
}
// @lc code=end

