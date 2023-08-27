// https://leetcode.cn/problems/longest-univalue-path
// 最长同值路径
package main

func longestUnivaluePath(root *TreeNode) (maxPath int) {
	if root == nil {
		return 0
	}
	path(root, &maxPath)
	return
}

func path(x *TreeNode, maxPath *int) int {
	if x == nil {
		return 0
	}
	left := path(x.Left, maxPath)
	right := path(x.Right, maxPath)
	leftLen, rightLen := 0, 0
	if x.Left != nil && x.Left.Val == x.Val {
		leftLen = left + 1
	}
	if x.Right != nil && x.Right.Val == x.Val {
		rightLen = right + 1
	}
	*maxPath = max(*maxPath, leftLen+rightLen)
	return max(leftLen, rightLen)
}
