package base

/*
快速排序（Quick Sort）

快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

算法描述

快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

从数列中挑出一个元素，称为 “基准”（pivot）；
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

左右指针法：
      左右指针法实现思路：在一段区间内我们有一个值key，从左边区间进行遍历，直到找到一个大于key的值就停下，然后再从右边找小于key的值，找到一个也停下来。我们将左右的值进行交换，这样左边那个大于key的值就被换到了右边，而右边那个比key小的值就被换到了左边。当左右两个指针相遇的时候就说明所有元素都与key做过了比较。然后再将左指针所在的元素赋值给key。此时按照上述方法进行递归实现[left, key]和[key+1, right]
 */
// 快排双指针法必须从右边往左边遍历，原因我已经想明白了

func QuickSort(slice []int, left, right int) {
	// 递归出口
	if left >= right {
		return
	}
	i, j, key := left, right, left
	for i < j {
		for i < j && slice[key] < slice[j] {
			j --
		}
		for i < j && slice[key] >= slice[i] {
			i ++
		}
		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[key], slice[i] = slice[i], slice[key]
	QuickSort(slice, left, i-1)
	QuickSort(slice, i+1, right)
}
