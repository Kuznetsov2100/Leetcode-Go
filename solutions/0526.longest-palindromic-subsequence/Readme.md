## 题目地址

https://leetcode-cn.com/problems/longest-palindromic-subsequence/

## 题目描述

```
给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。

 

示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2
一个可能的最长回文子序列为 "bb"。

 

提示：

1 <= s.length <= 1000
s 只包含小写英文字母
```

## 关键点

dp[i][j]的含义：在子串 s[i..j] 中，最长回文子序列的长度为dp[i][j]