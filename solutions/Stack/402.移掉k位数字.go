/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	var stack []byte // 单调递增栈,栈中存储的是num字符串中的字符
	var res string
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	// 如果k没有用完，截取前面len(stack)-k个元素，这相当于从stack中pop了k个元素
	// 然后使用TrimLeft去除前导零
	res = strings.TrimLeft(string(stack[:len(stack)-k]), "0")
	if res == "" {
		return "0"
	}
	return res
}

// func removeKdigits(num string, k int) string {
// 	var stack []int // 单调递增栈,栈中存储的是num字符串中的字符索引
// 	var res []byte
// 	for i := range num {
// 		for k > 0 && len(stack) > 0 && num[stack[len(stack)-1]] > num[i] {
// 			stack = stack[:len(stack)-1]
// 			k--
// 		}
// 		stack = append(stack, i)
// 	}

// 	// 如果k没用完且栈非空，继续出栈
// 	for k > 0 && len(stack) > 0 {
// 		stack = stack[:len(stack)-1]
// 		k--
// 	}
// 	for i := range stack {
// 		res = append(res, num[stack[i]])
// 	}

// 	// 跳过结果前缀可能存在的前导零
// 	j := 0
// 	for j < len(res) && res[j] == '0' {
// 		j++
// 	}
// 	ans := string(res[j:])
// 	if ans == "" {
// 		return "0"
// 	}
// 	return ans
// }
// @lc code=end

