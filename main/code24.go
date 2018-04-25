package main

/*
给定一个非负整数 num。 对于范围 0 ≤ i ≤ num 中的每个数字 i ，计算其二进制数中的1的数目并将它们作为数组返回。

示例：
比如给定 num = 5 ，应该返回 [0,1,1,2,1,2].

 */

func countBits(num int) []int {
	var slice []int
	for i := 0; i <= num; i++ {
		if i == 0 {
			slice = append(slice, 0)
		} else {
			if i%2 == 1 {
				slice = append(slice, slice[i-1]+1)
			} else {
				x := i / 2
				for x != 3 && x != 5 && x != 7 && x != 9 && x%2 == 0 {
					x /= 2
				}
				if x == 1 {
					slice = append(slice, 1)
				} else if x == 3 || x == 5 || x == 7 || x == 9 {
					slice = append(slice, slice[i-1]-1)
				} else {
					slice = append(slice, slice[i-1])
				}
			}
		}
	}
	return slice
}
