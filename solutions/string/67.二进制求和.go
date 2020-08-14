/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
import "strings"
import "strconv"
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	// 如果num1和num2长度不一样，则在长度较小的字符串前面添0，使得两个字符串长度一致
	var leadingzero strings.Builder
	for  i := 0; i < abs(m-n); i++ {
		leadingzero.WriteString("0")
	}
	if m > n {
		leadingzero.WriteString(b)
		b = leadingzero.String()
	} else if m < n {
		leadingzero.WriteString(a)
		a = leadingzero.String()
	}
	res := make([]int, len(a)+1) // 加法结果的位数最多为对齐后的字符串长度+1
	for i := len(a)-1; i >= 0; i-- {
		add := int((a[i]-'0')+(b[i] - '0'))
		p1 , p2 := i,i+1 // add 在res数组中的索引
		// 叠加到res上
		sum := add + res[p2]
		res[p2] = sum % 2
		res[p1] += sum / 2
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

