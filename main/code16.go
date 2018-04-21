package main

/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 */

func maxSubArray(nums []int) int {
	var maxSlice []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxSlice = append(maxSlice, appendMore(0, nums[i]))
		} else {
			maxSlice = append(maxSlice, appendMore(maxSlice[i-1], nums[i]))
		}
	}
	max := maxSlice[0]
	for _, x := range maxSlice {
		if max < x {
			max = x
		}
	}
	return max
}

func appendMore(preNumber, number int) int {
	if preNumber >= 0 {
		return preNumber + number
	}
	return number
}

func main() {
	var array []int
	array = append(array, -2, 1, -3, 4, -1, 2, 1, -5, 4)
	maxSubArray(array)
}



// 把上面的办法简化
func maxSubArrayGreat(a []int) int {
	max := -(1<<63)
	crt := 0

	for _, v := range a {
		if crt + v < v {
			crt = v
		} else {
			crt += v
		}

		if crt > max {
			max = crt
		}
	}

	return max
}