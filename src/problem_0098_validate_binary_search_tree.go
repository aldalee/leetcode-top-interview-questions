// https://leetcode.cn/problems/validate-binary-search-tree/
// 验证二叉搜索树
package main

func isValidBST(root *TreeNode) bool {
	var pre *int
	cur := root
	for cur != nil {
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
		if pre != nil && *pre >= cur.Val { // validate
			return false
		}
		pre = &cur.Val
		cur = cur.Right
	}
	return true
}
