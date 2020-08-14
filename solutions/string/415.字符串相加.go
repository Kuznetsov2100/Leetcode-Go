/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
import "strings"
import "strconv"
func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	// 如果num1和num2长度不一样，则在长度较小的字符串前面添0，使得两个字符串长度一致
	var leadingzero strings.Builder
	for  i := 0; i < abs(m-n); i++ {
		leadingzero.WriteString("0")
	}
	if m > n {
		leadingzero.WriteString(num2)
		num2 = leadingzero.String()
	} else if m < n {
		leadingzero.WriteString(num1)
		num1 = leadingzero.String()
	}
	res := make([]int, len(num1)+1) // 加法结果的位数最多为对齐后的字符串长度+1
	for i := len(num1)-1; i >= 0; i-- {
		add := int((num1[i]-'0')+(num2[i] - '0'))
		p1 , p2 := i,i+1 // add 在res数组中的索引
		// 叠加到res上
		sum := add + res[p2]
		res[p2] = sum % 10
		res[p1] += sum / 10
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

