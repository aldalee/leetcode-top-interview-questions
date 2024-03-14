// 使用priorityqueue库函数，LeetCode运行
// 我的发问帖：https://stackoverflow.com/questions/78158922/why-cannot-index-priorityqueue-newwith-value-of-type-funccomparator-utils-comp
package main

import pq "github.com/emirpasic/gods/queues/priorityqueue"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := pq.NewWith(func(a, b any) int {
		return a.(*ListNode).Val - b.(*ListNode).Val
	})
	for _, list := range lists {
		if list != nil {
			heap.Enqueue(list)
		}
	}
	if heap.Empty() {
		return nil
	}
	v, _ := heap.Dequeue()
	head := v.(*ListNode)
	cur, pre := head, head
	for !heap.Empty() {
		if cur.Next != nil {
			heap.Enqueue(cur.Next)
		}
		v, _ := heap.Dequeue()
		cur = v.(*ListNode)
		pre.Next = cur
		pre = cur
	}
	return head
}
