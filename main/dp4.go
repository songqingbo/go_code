package main

import "fmt"

/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
 */

func numTrees(n int) int {
	slice := getSingleNumberSlice(n)
	return slice[n]
}

func getSingleNumberSlice(n int) []int {
	slice := make([]int, n+1)
	if n == 1 {
		slice[n] = 1
		return slice
	}
	slice[1] = 1
	for i := 2; i <= n; i++ {
		for j := i; j > 1; j-- {
			slice[i] += slice[j-1]
		}
		slice[i] += slice[i-1]
	}
	return slice
}

func main() {
	fmt.Println(numTrees(4))
}
