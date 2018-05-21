package main

/*

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

例如，

给定数组 [1,1,1,2,2,3] , 和 k = 2，返回 [1,2]。

注意：

你可以假设给定的 k 总是合理的，1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
 */

func topKFrequent(nums []int, k int) []int {
	m := findMax(nums)
	array := make([]int, m)
	// O(N)
	for _, x := range nums {
		array[x] ++
	}

	var slice []int
	for i := k; i > 0; i-- {
		tempIndex := 0
		for index, x := range array {
			if array[tempIndex] < x {
				tempIndex = index
			}
		}
		array[tempIndex] = -1
		slice = append(slice, tempIndex)
	}
	return slice
}

// O(N)
func findMax(nums []int) int {
	m := nums[0]
	for _, x := range nums {
		if x > m {
			m = x
		}
	}
	return m
}
