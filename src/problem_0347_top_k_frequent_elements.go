// https://leetcode.cn/problems/top-k-frequent-elements/
// 前K个高频元素
package main

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

type T struct {
	num int
	cnt int
}

func topKFrequent(nums []int, k int) []int {
	cntmap := map[int]int{}
	for _, num := range nums {
		cntmap[num]++
	}
	heap := priorityqueue.NewWith[T](func(a, b T) int {
		return a.cnt - b.cnt
	})
	for num, cnt := range cntmap {
		if heap.Size() < k {
			heap.Enqueue(T{num, cnt})
		} else if t, _ := heap.Peek(); cnt > t.cnt {
			heap.Enqueue(T{num, cnt})
		}
		if heap.Size() > k {
			heap.Dequeue()
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		t, _ := heap.Dequeue()
		ans[i] = t.num
	}
	return ans
}
