// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 二叉树的中序遍历
package main

// based on morris traversal to inorder traversal
func inorderTraversal(root *TreeNode) (res []int) {
	cur, mostRight := root, root
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return
}
