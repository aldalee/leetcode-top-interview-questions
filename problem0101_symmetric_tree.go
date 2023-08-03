// https://leetcode.cn/problems/symmetric-tree/
// 对称二叉树
package main

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}
