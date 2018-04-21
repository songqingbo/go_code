package main

/*
你是一个专业的强盗，计划抢劫沿街的房屋。每间房都藏有一定的现金，阻止你抢劫他们的唯一的制约因素就是相邻的房屋有保安系统连接，如果两间相邻的房屋在同一晚上被闯入，它会自动联系警方。

给定一个代表每个房屋的金额的非负整数列表，确定你可以在没有提醒警方的情况下抢劫的最高金额。
 */

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return getMaxNumber(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return getMaxNumber(nums[0]+nums[2], nums[1])
	}
	var dp []int
	dp = append(dp, getMaxNumber(nums[0]+nums[2], nums[1]))
	dp = append(dp, nums[1])
	for i := 3; i < len(nums); i++ {
		dp = append(dp, max(dp[i-2]+nums[i], dp[i-1]))
	}
	max := dp[0]
	for _, x := range dp {
		if x > max {
			max = x
		}
	}
	return max
}

func getMaxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 此方案可以完成，测试用例59/69 然而超过了时间限制
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return getMaxNumber(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return getMaxNumber(nums[0]+nums[2], nums[1])
	}
	return getMaxNumber(nums[0]+rob2(nums[2:]), nums[1]+rob2(nums[3:]))
}

// 弱智行为。。
func robSB(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	length := len(nums)

	var maxI int
	var maxJ int
	for i := 0; i < length; i = i + 2 {
		maxI += nums[i]
	}
	for j := 1; j < length; j = j + 2 {
		maxJ += nums[j]
	}
	if maxI > maxJ {
		return maxI
	}
	return maxJ
}
