// https://leetcode.cn/problems/find-median-from-data-stream/
// 数据流的中位数
package main

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

type MedianFinder struct {
	minHeap *priorityqueue.Queue[int]
	maxHeap *priorityqueue.Queue[int]
}

func Constructor() MedianFinder {
	return MedianFinder{
		priorityqueue.NewWith(func(a, b int) int { return a - b }),
		priorityqueue.NewWith(func(a, b int) int { return b - a }),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxHeap.Empty() {
		mf.maxHeap.Enqueue(num)
	} else if top, _ := mf.maxHeap.Peek(); top >= num {
		mf.maxHeap.Enqueue(num)
	} else {
		mf.minHeap.Enqueue(num)
	}
	mf.balanceHeaps()
}

func (mf *MedianFinder) balanceHeaps() {
	if mf.maxHeap.Size()-mf.minHeap.Size() == 2 {
		top, _ := mf.maxHeap.Dequeue()
		mf.minHeap.Enqueue(top)
	} else if mf.minHeap.Size()-mf.maxHeap.Size() == 2 {
		top, _ := mf.minHeap.Dequeue()
		mf.maxHeap.Enqueue(top)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	n := mf.maxHeap.Size() - mf.minHeap.Size()
	maxTop, _ := mf.maxHeap.Peek()
	minTop, _ := mf.minHeap.Peek()
	if n == 0 {
		return float64(maxTop+minTop) / 2
	} else if n >= 1 {
		return float64(maxTop)
	} else {
		return float64(minTop)
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
