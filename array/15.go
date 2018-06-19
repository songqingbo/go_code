package array

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

func ThreeSum(nums []int) [][]int {
	m := make(map[int]int)
	for x := range nums {
		m[x] = 1
	}
	var s [][]int
	var slice [][]int
	for i := 0; i < len(nums); i++ {
		slice = twoSum(nums, 0-nums[i], m)
		if len(slice) > 0 {
			for j := 0; j < len(slice); j++ {
				slice[j] = append(slice[j], nums[i])
				s = append(s, slice[j])
			}
		}
	}
	return s
}

func twoSum(nums []int, target int, m map[int]int) ([][]int) {
	var slice [][]int
	for k := range m {
		if m[target-k] == 1 {
			tempSlice := []int{k, target - k}
			m[target-k] = 0
			slice = append(slice, tempSlice)
		}
	}
	return slice
}
