package main

/*
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 */

func productExceptSelf(nums []int) []int {
	length := len(nums)
	// 特殊处理必不可少
	if length <= 1 {
		return nums
	}
	slice := make([]int, length)
	prefix := 1
	for i := 0; i < length; i++ {
		slice[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := length - 1; i >= 0; i-- {
		slice[i] *= suffix
		suffix *= nums[i]
	}
	return slice
}

/*
思路其实很简单。对于数组中的下表为index的数字，只需要计算在此数字之前的所有数字的乘积乘以在此数字之后的所有数字的乘积即可。
因此，可以遍历数组两次，第一次，计算所有在当前数字之前的数字的乘积。第二次，计算所有在当前数字之后的数字的乘积。只不过遍历方向分别是从前向后遍历和从后向前遍历。
 */
