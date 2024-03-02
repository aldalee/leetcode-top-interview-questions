// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 二叉树的前序遍历
package main

func preorderTraversal(root *TreeNode) (res []int) {
	cur := root
	for cur != nil {
		mostRight := cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				res = append(res, cur.Val)
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			res = append(res, cur.Val)
		}
		cur = cur.Right
	}
	return
}
