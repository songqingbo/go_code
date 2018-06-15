package main

/*
给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n(10的n次方)。

示例:
给定 n = 2，返回 91。（答案应该是除[11,22,33,44,55,66,77,88,99]外，0 ≤ x < 100 间的所有数字）
 */

func countNumbersWithUniqueDigits(n int) int {
	slice := make([]int, n+1)
	if n == 0 {
		return 1
	}
	slice[0] = 1
	for i := 1; i <= n; i++ {
		slice[i] = mulNumber(i) + slice[i-1]
	}
	return slice[n]
}

func mulNumber(n int) int {
	number := 9
	if n == 1 {
		return number
	}
	for i := 2; i <= n; i++ {
		number = number * (11 - i)
	}
	return number
}

/*
核心思想是按照数字的位数划分，以位数为思考基石，不要陷入数字值的误区
而且题目出得已经在提示答题者了，告诉是10的n次方，也应该往位数方向想
 */

 /*
 This problem can be seen as a blanks filling problem. The first blank can be taken by 1~9. The second blank has 10-1=9 choices(0~9 - one numbre has been chosen for the first blank). The third one has 10-2=8 choices.
dp[0] = 1
dp[1] = 9x9 + dp[0]
dp[2] = 9x9x8 + dp[1]
dp[3] = 9x9x8x7 + dp[2]
...
  */