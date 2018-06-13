package main

import "fmt"

/*
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
 */

func minimumTotal(triangle [][]int) int {
	// layers of triangle
	m := len(triangle)
	slice := make([][]int, m)
	for i := 0; i < m; i++ {
		tempSlice := make([]int, m)
		slice[i] = tempSlice
		if i == 0 {
			slice[0][0] = triangle[0][0]
			continue
		}
		slice[i][0] = triangle[i][0] + slice[i-1][0]
		slice[i][i] = triangle[i][i] + slice[i-1][i-1]
	}
	for i := 2; i < m; i++ {
		for j := 1; j < i; j++ {
			slice[i][j] = triangle[i][j] + getMin(slice[i-1][j], slice[i-1][j-1])
		}
	}
	min := 999999999999
	for i := 0; i < m; i++ {
		if slice[m-1][i] < min {
			min = slice[m-1][i]
		}
	}
	return min
}

func getMin(n, m int) int {
	if n > m {
		return m
	}
	return n
}

func main() {
	slice := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(slice))
}
