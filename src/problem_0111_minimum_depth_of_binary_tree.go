// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 二叉树的最小深度
package main

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := root
	height := 0
	res := math.MaxInt
	for cur != nil {
		if cur.Left == nil { // only once arrived
			height++
			if cur.Right == nil { // the original leaf node
				res = min(height, res)
			}
			cur = cur.Right
		} else {
			count := 1
			rightmost := cur.Left
			for rightmost.Right != nil && rightmost.Right != cur {
				count++
				rightmost = rightmost.Right
			}
			if rightmost.Right == nil { // first arrived
				height++
				rightmost.Right = cur
				cur = cur.Left
			} else { // second arrived
				if rightmost.Left == nil {
					res = min(height, res)
				}
				height -= count
				rightmost.Right = nil
				cur = cur.Right
			}
		}
	}
	return res
}
