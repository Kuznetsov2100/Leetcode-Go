/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
import "sort"
import "strings"
func frequencySort(s string) string {	
	type foo struct {
		alpha byte
		frequency int
	}
	var arr []foo
	var res strings.Builder
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	for k, v := range m {
		arr = append(arr, foo{alpha: k, frequency: v})
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].frequency > arr[j].frequency })
	for i := range arr {
		for j := 0; j < arr[i].frequency; j++ {
			res.WriteByte(arr[i].alpha)
		}
	}
	return res.String()
}
// @lc code=end

