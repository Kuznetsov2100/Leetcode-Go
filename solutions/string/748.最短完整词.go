/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短完整词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	var res string
	minLength := 16
	licensePlate = simplify(licensePlate)
	target := count(licensePlate)
	for _, word := range words {
		n := len(word)
		if n < minLength && isValid(count(word), target) {
			minLength = n
			res = word
		}
	}
	return res
}

func simplify(s string) string {
	var b bytes.Buffer
	for i := range s {
		if 'a' <= s[i] && s[i] <= 'z' || 'A' <= s[i] && s[i] <= 'Z' {
			b.WriteByte(s[i] | ' ')
		}
	}
	return b.String()
}

func isValid(arr1, arr2 []int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] < arr2[i] {
			return false
		}
	}
	return true
}

func count(word string) []int {
	m := make([]int, 26)
	for i := range word {
		m[word[i]-'a']++
	}
	return m
}
// @lc code=end

