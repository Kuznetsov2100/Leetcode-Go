/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(input string) []int {
// 对于一个形如 x op y（op为运算符，x和y为数）的运算表达式而言，
// 它的结果组合取决于x和y的结果组合数，而x和y又可以写成形如x op y的算式。
// 因此，该问题的子问题就是x op y中的x和y：以运算符分隔的左右两侧算式解。
// 然后我们来进行 分治算法三步走：
// 分解：按运算符分成左右两部分，分别求解
// 解决：实现一个递归函数，输入算式，返回算式解
// 合并：根据运算符合并左右两部分的解，得出最终解

	if !strings.ContainsAny(input, "+-*") {
		num, _ := strconv.Atoi(input)
		return []int{num}
	}
	var res []int
	for i := range input {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			for _, x := range diffWaysToCompute(input[:i]) {
				for _, y := range diffWaysToCompute(input[i+1:]) {
					res = append(res, compute(x, y, input[i]))
				}
			}
		}
	}
	return res
}

func compute(a, b int, op byte) int {
	var res int
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	}
	return res
}
// @lc code=end

