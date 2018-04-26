package main

/*
给定一个非负整数 num。 对于范围 0 ≤ i ≤ num 中的每个数字 i ，计算其二进制数中的1的数目并将它们作为数组返回。

示例：
比如给定 num = 5 ，应该返回 [0,1,1,2,1,2].

 */

func countBits(num int) []int {
	var slice []int
	for i := 0; i <= num; i++ {
		tempSlice := intToBinaryString(i)
		slice = append(slice, get1N(tempSlice))
	}
	return slice
}

func intToBinaryString(number int) []rune {
	var tempSlice []rune
	for number > 0 {
		tempSlice = append(tempSlice, rune(number%2))
		number = number / 2
	}
	var slice []rune
	for i := len(tempSlice) - 1; i > -1; i-- {
		slice = append(slice, tempSlice[i])
	}
	return slice
}

func get1N(slice []rune) int {
	n := 0
	for _, x := range slice {
		if x == '1' {
			n ++
		}
	}
	return n
}
