// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree
// 将有序数组转换为二叉搜索树
package main

func sortedArrayToBST(nums []int) *TreeNode {
	var f func(l, r int) *TreeNode
	f = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		if l == r {
			return &TreeNode{Val: nums[r]}
		}
		m := ((r - l) >> 1) + l
		root := &TreeNode{Val: nums[m]}
		root.Left = f(l, m-1)
		root.Right = f(m+1, r)
		return root
	}
	return f(0, len(nums)-1)
}
