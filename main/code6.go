package main

/*
给定一个数组 nums, 编写一个函数将所有 0 移动到它的末尾，同时保持非零元素的相对顺序。

例如， 定义 nums = [0, 1, 0, 3, 12]，调用函数之后， nums 应为 [1, 3, 12, 0, 0]。

注意事项:

必须在原数组上操作，不要为一个新数组分配额外空间。
尽量减少操作总数。
 */

// 能力一般水平有限
func moveZeroes2(nums []int) {
	var array [len(nums)]int
	i := 0
	for _, x := range nums {
		if x != 0 {
			array[i] = x
			i ++
		}
	}
	for index, x := range array {
		nums[index] = x
	}
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j ++
		}
	}
}
