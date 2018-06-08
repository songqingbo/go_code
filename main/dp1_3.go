package main

// 动态规划解法，前两个其实就是暴力解法
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	var slice [10000][10000]bool
	slice[0][0] = true
	maxLength := 0
	var str string
	for i := 0; i < length-1; i++ {
		for j := i; j > 0; j-- {
			slice[i][j] = s[i] == s[j] && (i-j < 3 || slice[i+1][j-1])
			if slice[i][j] && (maxLength < j-i+1) {
				str = s[i:j+1]
			}
		}
	}
	return str
}
