package array

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。画 n 条垂直线，使得垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

注意：你不能倾斜容器，n 至少是2。
 */

func maxArea(height []int) int {
	length := len(height)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			temp := (j - i) * getMinN(height[i], height[j])
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

func getMinN(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// O（n）
func maxArea2(height []int) int {
	length := len(height)
	max := 0
	begin := 0
	end := length - 1
	for begin < end {
		tempMax := (end - begin) * getMinN(height[begin], height[end])
		if max < tempMax {
			max = tempMax
		}
		// move
		if height[begin] < height[end] {
			begin ++
		} else {
			end --
		}
	}
	return max
}
