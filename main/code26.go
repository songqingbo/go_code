package main

/*

给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */

 // 递归
func inorderTraversal2(root *TreeNode) []int {
	var slice []int
	helper(root, &slice)
	return slice
}

func helper(root *TreeNode, slice *[]int) {
	if root != nil {
		helper(root.Left, slice)
		*slice = append(*slice, root.Val)
		helper(root.Right, slice)
	}
}
