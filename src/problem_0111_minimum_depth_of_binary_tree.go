// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 二叉树的最小深度
package main

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth, curDepth := math.MaxInt, 0
	for cur := root; cur != nil; {
		if cur.Left == nil { // only once arrived
			curDepth++
			if cur.Right == nil { // the original leaf node
				minDepth = min(curDepth, minDepth)
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
				curDepth++
				rightmost.Right = cur
				cur = cur.Left
			} else { // second arrived
				if rightmost.Left == nil {
					minDepth = min(curDepth, minDepth)
				}
				curDepth -= count
				rightmost.Right = nil
				cur = cur.Right
			}
		}
	}
	return minDepth
}
