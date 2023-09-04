// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
// 二叉树的锯齿形层序遍历
package main

import "container/list"

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	deque := list.New()
	deque.PushBack(root)
	for deque.Len() > 0 {
		var level []int
		for i := deque.Len(); i > 0; i-- {
			cur := deque.Remove(deque.Front()).(*TreeNode)
			if len(ans)%2 == 0 {
				level = append(level, cur.Val)
			} else {
				level = append([]int{cur.Val}, level...)
			}
			if cur.Left != nil {
				deque.PushBack(cur.Left)
			}
			if cur.Right != nil {
				deque.PushBack(cur.Right)
			}
		}
		ans = append(ans, level)
	}
	return
}
