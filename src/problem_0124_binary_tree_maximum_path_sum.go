// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 二叉树中的最大路径和
package main

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31
	maxGain(root, &maxSum)
	return maxSum
}

// 从当前节点出发向下走的最大路径和
func maxGain(x *TreeNode, maxSum *int) int {
	if x == nil {
		return 0
	}
	leftGain := max(maxGain(x.Left, maxSum), 0)
	rightGain := max(maxGain(x.Right, maxSum), 0)
	*maxSum = max(*maxSum, x.Val+leftGain+rightGain)
	return x.Val + max(leftGain, rightGain)
}
