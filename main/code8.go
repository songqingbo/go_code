package main

/*
给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

例如：

输入: 二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
 */

// 遍历处理OK
func convertBST2(root *TreeNode) *TreeNode {
	var array []int
	getIntArray(&array, root)
	setIntValue(array, root)
	return root
}

func getIntArray(array *[]int, root *TreeNode) {
	if root != nil {
		*array = append(*array, root.Val)
		getIntArray(array, root.Left)
		getIntArray(array, root.Right)
	}
}

func setIntValue(array []int, root *TreeNode) {
	if root != nil {
		n := getAdder(array, root.Val)
		root.Val += n
		setIntValue(array, root.Left)
		setIntValue(array, root.Right)
	}
}

func getAdder(array []int, x int) int {
	n := 0
	for _, y := range array {
		if y > x {
			n += y
		}
	}
	return n
}

func main() {
	left := TreeNode{Val: 2, Right: nil, Left: nil}
	right := TreeNode{Val: 13, Right: nil, Left: nil}
	root := TreeNode{Val: 5, Right: &right, Left: &left}
	convertBST(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// 递归
var sum int
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Right)
	root.Val += sum
	sum = root.Val
	traverse(root.Left)
}