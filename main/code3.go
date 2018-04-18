package main

/*

给定一个整数数组，除了某个元素外其余元素均出现两次。请找出这个只出现一次的元素。



备注：

1.你的算法应该是一个线性时间复杂度
2.你可以不用额外空间来实现它吗？
 */

// 异或
func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}

// map
func singleNumber2(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if mnum, ok := m[num]; !ok {
			m[num] = 1
		} else if mnum == 1 {
			delete(m, num)
		}
	}
	// only 1 element should be left
	for res := range m {
		return res
	}
	return 0 // never reach here
}
