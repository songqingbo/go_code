package main

import "fmt"

/*

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
 */

func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	m := stringSliceToMap(wordDict)
	slice := make([]bool, length+1)
	slice[0] = true
	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if slice[i] {
				break
			}
			slice[i] = slice[j] && m[s[j:i]]
		}
	}
	return slice[length]
}

func stringSliceToMap(slice []string) map[string]bool {
	length := len(slice)
	m := make(map[string]bool)
	for i := 0; i < length; i++ {
		m[slice[i]] = true
	}
	return m
}

func main() {
	fmt.Println(wordBreak("applepie", []string{"pie", "pear", "apple", "peach"}))
}