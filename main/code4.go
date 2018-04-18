package main

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶节点的最长路径上的节点数。

案例：
给出二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回最大深度为 3 。
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}

	return 1 + maxNumber(maxDepth(root.Right), maxDepth(root.Left))
}

func maxNumber(x, y int) int {
	if x > y {
		return x
	}
	return y
}
