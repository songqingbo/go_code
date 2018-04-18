package main

/*
翻转一棵二叉树。

     4
   /   \
  2     7
 / \   / \
1   3 6   9
转换为：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您无法在白板上写出翻转二叉树这道题，这太糟糕了。
 */

// 居然一次通过了。。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tempNode := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(tempNode)
	return root
}

// 帮助理解
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	tempNode := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(tempNode)
	return root
}

// 最简洁
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
