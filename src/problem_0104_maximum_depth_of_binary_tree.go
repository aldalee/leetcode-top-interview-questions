// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 二叉树的最大深度
package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth, curDepth := 0, 0
	for cur := root; cur != nil; {
		if cur.Left == nil {
			curDepth++
			cur = cur.Right
		} else {
			rightmost := cur.Left
			count := 1
			for rightmost.Right != nil && rightmost.Right != cur {
				count++
				rightmost = rightmost.Right
			}
			if rightmost.Right == nil {
				curDepth++
				rightmost.Right = cur
				cur = cur.Left
			} else {
				maxDepth = max(maxDepth, curDepth)
				curDepth -= count
				rightmost.Right = nil
				cur = cur.Right
			}
		}
	}
	return max(maxDepth, curDepth)
}
