package main

/*
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
 */


/*
public int findUnsortedSubarray(int[] A) {
    int n = A.length, beg = -1, end = -2, min = A[n-1], max = A[0];
    for (int i=1;i<n;i++) {
      max = Math.max(max, A[i]);
      min = Math.min(min, A[n-1-i]);
      if (A[i] < max) end = i;
      if (A[n-1-i] > min) beg = n-1-i;
    }
    return end - beg + 1;
}
 */

 // 此方案过了50%的测试用例
func findUnsortedSubarray2(nums []int) int {
	numsLength := len(nums)
	if numsLength <= 1 {
		return 0
	}
	if numsLength == 2 {
		if nums[0] <= nums[1] {
			return 0
		} else {
			return 2
		}
	}
	min, max := getMinAndMax(nums)
	start := 0
	end := len(nums) - 1
	for start < end {
		if nums[start] == nums[end] {
			start = end
		} else if nums[start] == min && nums[end] == max {
			start ++
			min, max = getMinAndMax(nums[start:end])
			end --
		} else if nums[start] > min || nums[end] < max {
			if nums[start] == min {
				start ++
			}
			if nums[end] == max {
				end --
			}
			break
		} else {
			break
		}
	}
	return end - start + 1
}

func getMinAndMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		return -1, -1
	}
	min, max = nums[0], nums[0]
	for _, x := range nums {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return min, max
}
