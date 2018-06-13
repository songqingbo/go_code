package main

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	slice := make([]int, length+1)
	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			slice[j] = minInt(slice[j], slice[j+1]) + triangle[i][j]
		}
	}
	return slice[0]
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
