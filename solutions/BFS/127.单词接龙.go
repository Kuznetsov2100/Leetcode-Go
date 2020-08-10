/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
// 最重要的步骤是找出只差一个字母的多个单词，即从当前单词一步能得到哪些单词。
// 为了快速的找到这些相邻节点，我们对给定的 wordList 做一个预处理，将单词中的某个字母用 * 代替
// 这个预处理帮我们构造了一个单词变换的通用状态。

// 我们使用一个Map来存储，键：通用状态；值：能得到这个通用状态的所有单词
// 例如：Dog ----> D*g <---- Dig，Dog 和 Dig 都指向了一个通用状态 D*g。即键D*g对应的值为[Dog,Dig]

// 这步预处理找出了单词表中所有单词改变某个字母后的通用状态，并帮助我们更方便也更快的找到相邻节点。
// 否则，对于每个单词我们需要遍历整个字母表查看是否存在一个单词与它相差一个字母，这将花费很多时间。
// 预处理操作在广度优先搜索之前高效的建立了邻接表。

// 例如，在广搜时我们需要访问 Dug 的所有邻接点，也就是想知道由Dug能变换到字典中哪些单词
// 我们可以先生成 Dug 的所有通用状态：以通用状态为键去获取值就可得到结果
// 	* Dug => *ug
// Dug => D*g
// Dug => Du*
// 第二个变换 D*g 可以同时映射到 Dog 或者 Dig，因为他们都有相同的通用状态。拥有相同的通用状态意味着两个单词只相差一个字母，他们的节点是相连的。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	transform := make(map[string][]int)
	for index, word := range wordList {
		for i := range word {
			k := word[:i] + "*" + word[i+1:]
			if _, ok := transform[k]; !ok {
				transform[k] = []int{}
			}
			transform[k] = append(transform[k], index)
		}
	}
	// 开始词作为第一个节点加入队列,深度count是1，标记其已访问
	count := 1
	visited := make([]bool, len(wordList))
	queue := []string{beginWord}
	
	for len(queue) > 0 {
		count++
		var nextQueue[]string
		for _, word := range queue {
			for i := range word {
				key := word[:i] + "*" + word[i+1:]
				// 拿到这个通配词映射的单词集合(也就是从当前单词一次转换能得到哪些单词)
				for _, index := range transform[key] {
					if  wordList[index] == endWord {
						return count
					}
					if !visited[index] { // 只访问未访问过的节点
						visited[index] = true
						nextQueue = append(nextQueue, wordList[index]) // 将单词加入到下一轮迭代队列中
					}
				}
			}
		}
		queue = nextQueue
	}
	return 0 // 转换失败
}
// @lc code=end

