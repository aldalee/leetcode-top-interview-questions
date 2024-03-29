// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
// 二叉搜索树中第K小的元素
package main

func kthSmallest(root *TreeNode, k int) int {
	for cur := root; cur != nil; {
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
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}
