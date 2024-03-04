// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 二叉树的最近公共祖先
package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	nodesInLeft := lowestCommonAncestor(root.Left, p, q)
	nodesInRight := lowestCommonAncestor(root.Right, p, q)
	if nodesInLeft != nil && nodesInRight != nil {
		return root
	} else if nodesInLeft == nil {
		return nodesInRight
	}
	return nodesInLeft
}
