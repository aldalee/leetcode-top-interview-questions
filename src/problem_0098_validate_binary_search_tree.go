// https://leetcode.cn/problems/validate-binary-search-tree/
// 验证二叉搜索树
package main

func isValidBST(root *TreeNode) bool {
	var pre *int
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
		if pre != nil && *pre >= cur.Val { // validate
			return false
		}
		pre = &cur.Val
		cur = cur.Right
	}
	return true
}
