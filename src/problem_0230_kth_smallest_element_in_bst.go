// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
// 二叉搜索树中第K小的元素
package main

func kthSmallest(root *TreeNode, k int) int {
	cur := root
	for cur != nil {
		mostRight := cur.Left
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
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}
