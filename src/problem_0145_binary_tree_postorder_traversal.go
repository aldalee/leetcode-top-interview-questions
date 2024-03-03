// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 二叉树的后序遍历
package main

func postorderTraversal(root *TreeNode) (res []int) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else {
			mostRight := cur.Left
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
			} else {
				mostRight.Right = nil
				collectEdge(cur.Left, &res)
				cur = cur.Right
			}
		}
	}
	collectEdge(root, &res)
	return
}

// collect right boundary
func collectEdge(root *TreeNode, res *[]int) {
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
