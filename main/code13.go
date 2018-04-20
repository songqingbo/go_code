package main

/*

假设你正在爬楼梯。需要 n 步你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 步 + 1 步
2.  2 步
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 步 + 1 步 + 1 步
2.  1 步 + 2 步
3.  2 步 + 1 步
 */

// 超时了
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 递归转循环
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	total, one, two := 2, 1, 2
	for i := 2; i < n; i++ {
		total = one + two
		one = two
		two = total
	}
	return total
}

// 递归 动态规划
func climbStairs3(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	var slice [] int
	slice = append(slice, 1, 1)
	for i := 2; i <= n; i++ {
		slice = append(slice, slice[i-1]+slice[i-2])
	}
	return slice[n]
}
