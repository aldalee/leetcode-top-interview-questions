// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 二叉树的中序遍历
package main

// based on morris traversal to inorder traversal
func inorderTraversal(root *TreeNode) (res []int) {
	cur := root
	for cur != nil {
		rightmost := cur.Left
		if rightmost != nil {
			for rightmost.Right != nil && rightmost.Right != cur {
				rightmost = rightmost.Right
			}
			if rightmost.Right == nil {
				rightmost.Right = cur
				cur = cur.Left
				continue
			} else {
				rightmost.Right = nil
			}
		}
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return
}
