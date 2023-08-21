// https://leetcode.cn/problems/merge-k-sorted-lists/
// 合并K个升序链表
package main

import "container/heap"

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(v any)        { *pq = append(*pq, v.(*ListNode)) }
func (pq *PriorityQueue) Pop() any          { a := *pq; v := a[len(a)-1]; *pq = a[:len(a)-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}
	if pq.Len() == 0 {
		return nil
	}
	head := heap.Pop(pq).(*ListNode)
	cur := head
	pre := head
	for pq.Len() > 0 {
		if cur.Next != nil {
			heap.Push(pq, cur.Next)
		}
		cur = heap.Pop(pq).(*ListNode)
		pre.Next = cur
		pre = cur
	}
	return head
}
