// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 二叉树的层序遍历
package main

import "container/list"

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var level []int
		for i := queue.Len(); i > 0; i-- {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
		ans = append(ans, level)
	}
	return
}
