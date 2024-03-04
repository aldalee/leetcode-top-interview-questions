// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 二叉树的前序遍历
package main

func preorderTraversal(root *TreeNode) (res []int) {
	for cur := root; cur != nil; {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			rightmost := cur.Left
			for rightmost.Right != nil && rightmost.Right != cur {
				rightmost = rightmost.Right
			}
			if rightmost.Right == nil {
				rightmost.Right = cur
				res = append(res, cur.Val)
				cur = cur.Left
			} else {
				rightmost.Right = nil
				cur = cur.Right
			}
		}
	}
	return
}
