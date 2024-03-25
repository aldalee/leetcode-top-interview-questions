// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
// 有序矩阵中第K小的元素
package main

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

type node struct {
	val int
	row int
	col int
}

func kthSmallest(matrix [][]int, k int) int {
	heap := priorityqueue.NewWith[*node](func(a, b *node) int {
		return a.val - b.val
	})
	heap.Enqueue(&node{matrix[0][0], 0, 0})
	var ans *node
	for !heap.Empty() && k != 0 {
		ans, _ = heap.Dequeue()
		k--
		row, col := ans.row, ans.col
		if row+1 < len(matrix) && col == 0 {
			heap.Enqueue(&node{matrix[row+1][col], row + 1, col})
		}
		if col+1 < len(matrix) {
			heap.Enqueue(&node{matrix[row][col+1], row, col + 1})
		}
	}
	return ans.val
}
