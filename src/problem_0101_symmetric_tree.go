// https://leetcode.cn/problems/symmetric-tree/
// 对称二叉树
package main

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && // validate the value
		check(p.Left, q.Right) && // validate left subtree
		check(p.Right, q.Left) // validate right subtree
}
