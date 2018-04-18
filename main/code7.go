package main

/*
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, n] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:

输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]
 */

// 借助额外空间
func findDisappearedNumbers2(nums []int) []int {
	var array [4]int
	for _, x := range nums {
		array[x] ++
	}
	var arraySlice []int
	for index, x := range array {
		if x == 0 {
			arraySlice = append(arraySlice, index)
		}
	}
	return arraySlice
}

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	nums = append(nums, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	slice := nums[n:]
	for _, x := range nums {
		slice[x] = 1
	}
	for index, x := range slice {
		if x == 0 {
			slice = append(slice, index)
		}
	}
	return slice[len(nums):]
}

// 改成负数的解法
// O（n）时间复杂度排序