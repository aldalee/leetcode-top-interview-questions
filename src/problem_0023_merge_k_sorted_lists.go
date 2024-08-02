// https://leetcode.cn/problems/merge-k-sorted-lists/
// 合并K个升序链表
package main

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

func mergeKLists(lists []*ListNode) *ListNode {
	heap := priorityqueue.NewWith[*ListNode](func(a, b *ListNode) int {
		return a.Val - b.Val
	})
	for _, list := range lists {
		if list != nil {
			heap.Enqueue(list)
		}
	}
	if heap.Empty() {
		return nil
	}
	head, _ := heap.Dequeue()
	cur, pre := head, head
	for !heap.Empty() {
		if cur.Next != nil {
			heap.Enqueue(cur.Next)
		}
		cur, _ = heap.Dequeue()
		pre.Next = cur
		pre = cur
	}
	return head
}
