/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
import "strconv"
func calculate(s string) int {	
	var M,Q []string
	precedence := map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	operand := 0
	for i := range s {	
		if s[i] >= '0' && s[i] <= '9' {
			if i == 0 && s[i] == '0' {
				Q = append(Q, "0")
			} else if i > 0 && s[i] == '0' && !isNumber(string(s[i-1])) {
				Q = append(Q, "0")
			} else {
				operand = 10 * operand + int(s[i]- '0')
			}
		} else if s[i] == '(' {
			if operand != 0 {
				Q = append(Q, strconv.Itoa(operand))
				operand = 0
			}
			M = append(M, string(s[i]))
		} else if s[i] == ')' {
			if operand != 0 {
				Q = append(Q, strconv.Itoa(operand))
				operand = 0
			}
			for len(M) > 0 && M[len(M)-1] != "(" {
				Q = append(Q, M[len(M)-1])
				M = M[:len(M)-1]		
			}
			M = M[:len(M)-1]
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			if operand != 0 {
				Q = append(Q, strconv.Itoa(operand))
				operand = 0
			}
			for len(M) > 0 && precedence[M[len(M)-1]] >= precedence[string(s[i])]  {
				Q = append(Q, M[len(M)-1])
				M = M[:len(M)-1]		
			}
			M = append(M,string(s[i]))
		} 
		if i == len(s)-1 {	
			if operand != 0 {
				Q = append(Q, strconv.Itoa(operand))
				operand = 0
			}
		}
	}
	if e := Q[len(Q)-1]; len(M) == 0 && isNumber(e) {
		num, _ := strconv.Atoi(e)
		return num
	}
	for len(M) > 0 {
		Q = append(Q, M[len(M)-1])
		M = M[:len(M)-1]
	}
	return evalRPN(Q)
}


func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func evalRPN(tokens []string) int {
	var stack []int
	for i := range tokens {
		element := tokens[i]
		if element != "+" && element != "-" && element != "*" && element != "/" {
			num, _ := strconv.Atoi(element)
			stack = append(stack, num)			
		} else {
			numA := stack[len(stack)-1]
			numB := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if element == "+" {	
				stack = append(stack, numB+numA)
			} else if element == "-" {
				stack = append(stack, numB-numA)
			} else if element == "*" {
				stack = append(stack, numB*numA)
			} else if element == "/" {
				stack = append(stack, numB/numA)
			}
		}
	}
	return stack[0]
}
// @lc code=end

