/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
import "sort"
func topKFrequent(words []string, k int) []string {
	type foo struct {
		word string
		frequency int
	}
	var arr []foo
	var res []string
	m := make(map[string]int)
	for i := range words {
		m[words[i]]++
	}
	for k, v := range m {
		arr = append(arr, foo{word: k, frequency: v})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].frequency > arr[j].frequency {
			return true
		} else if arr[i].frequency == arr[j].frequency {
			if arr[i].word < arr[j].word {
				return true
			}
		}
		return false
	})
	for i := 0; i < k; i++ {
		res = append(res,arr[i].word)
	}
	return res
}
// @lc code=end

