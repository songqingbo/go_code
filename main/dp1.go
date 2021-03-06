package main

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

---

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
 */

// 82 / 94 个通过测试用例
func longestPalindrome1(s string) string {
	length := len(s)
	slice := make([]string, length)
	for i := 0; i < length; i++ {
		slice[i] = getMaxLengthString1(s[i:])
	}
	maxLengthString := ""
	for _, x := range slice {
		if len(x) > len(maxLengthString) {
			maxLengthString = x
		}
	}
	return maxLengthString
}

func isPalindromic1(s string) bool {
	if len(s) <= 1 {
		return true
	}
	flag := true
	i := 0
	j := len(s) - 1

	for i < j {
		flag = flag && (s[i:i+1] == s[j:j+1])
		i++
		j--
	}
	return flag
}

func getMaxLengthString1(s string) string {
	length := len(s)
	maxLength := 1
	if length <= 1 {
		return s
	}
	for i := length; i > 1; i-- {
		if isPalindromic1(s[0:i]) && maxLength < len(s[0:i]) {
			maxLength = len(s[0:i])
		}
	}
	return s[:maxLength]
}
