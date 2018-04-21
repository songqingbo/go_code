package main

/*

给定一个二叉树，二叉树的每个节点含有一个整数。

找出路径和等于给定数的路径总数。

路径不需要从根节点开始，也不需要在叶节点结束，当路径方向必须是向下的（只从父节点到子节点）。

二叉树不超过1000个节点，节点的整数值的范围是[-1000000,1000000]。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3. 和等于8的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
 */

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathNumberFrom(root, sum) + pathNumberFrom(root.Right, sum) + pathNumberFrom(root.Left, sum)
}

func pathNumberFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var n int
	if root.Val == sum {
		n = 1
	}
	return n + pathNumberFrom(root.Left, sum-root.Val) + pathNumberFrom(root.Right, sum-root.Val)
}
