// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 二叉树的最大深度
package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
