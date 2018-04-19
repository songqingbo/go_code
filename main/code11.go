package main

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
 */
// 还是差点
func getValue(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	leftTreeHigh := getHigh(root.Left)
	rightTreeHigh := getHigh(root.Right)
	if root.Right == nil || root.Left == nil {
		return leftTreeHigh + rightTreeHigh + 1
	}
	return leftTreeHigh + rightTreeHigh + 2
}

func getHigh(root *TreeNode) int {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return 0
	}
	return 1 + getMax(getHigh(root.Left), getHigh(root.Right))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 极致递归

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)), height(root.Left)+height(root.Right))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}
