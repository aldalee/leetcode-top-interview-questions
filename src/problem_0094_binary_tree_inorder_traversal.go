// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 二叉树的中序遍历
package main

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return
}
