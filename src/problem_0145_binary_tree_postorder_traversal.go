// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 二叉树的后序遍历
package main

func postorderTraversal(root *TreeNode) (res []int) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else {
			rightmost := cur.Left
			for rightmost.Right != nil && rightmost.Right != cur {
				rightmost = rightmost.Right
			}
			if rightmost.Right == nil {
				rightmost.Right = cur
				cur = cur.Left
			} else {
				rightmost.Right = nil
				collectRightmostEdge(cur.Left, &res)
				cur = cur.Right
			}
		}
	}
	collectRightmostEdge(root, &res)
	return
}

// collect right boundary
func collectRightmostEdge(root *TreeNode, res *[]int) {
	tail := reverseEdge(root)
	for cur := tail; cur != nil; cur = cur.Right {
		*res = append(*res, cur.Val)
	}
	reverseEdge(tail)
}

func reverseEdge(root *TreeNode) *TreeNode {
	var pre *TreeNode
	for cur := root; cur != nil; {
		next := cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
	return pre
}
