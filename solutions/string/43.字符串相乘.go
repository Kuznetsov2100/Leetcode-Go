/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
import "strings"
import "strconv"
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]int, m+n) // 结果最多为m+n位数
	// 从个位数开始逐位相乘
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			mul := int((num1[i]-'0')*(num2[j] - '0'))
			// 乘积在res数组中对应的索引位置
			p1 ,p2 := i+j, i+j+1
			// 叠加到res上
			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	// 跳过结果前缀可能存在的0(未使用的位)
	k := 0
	for k < len(res) && res[k] == 0 {
		k++
	}
	// 将计算结算转为string
	var str strings.Builder
	for ;k < len(res); k++ {
		str.WriteString(strconv.Itoa(res[k]))
	}
	if str.Len() == 0 {
		return "0"
	}
	return str.String()
}
// @lc code=end

