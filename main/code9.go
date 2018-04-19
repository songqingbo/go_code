package main

/*
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且数组中的众数永远存在。
 */

func majorityElement(nums []int) int {
	valueNumberMap := make(map[int]int)
	for _, x := range nums {
		valueNumberMap[x] += 1
	}
	for k, v := range valueNumberMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}
