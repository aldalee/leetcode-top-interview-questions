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

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Empty() {
		this.maxHeap.Enqueue(num)
	} else if top, _ := this.maxHeap.Peek(); top >= num {
		this.maxHeap.Enqueue(num)
	} else {
		this.minHeap.Enqueue(num)
	}
	this.balanceHeaps()
}

func (this *MedianFinder) balanceHeaps() {
	if this.maxHeap.Size()-this.minHeap.Size() == 2 {
		top, _ := this.maxHeap.Dequeue()
		this.minHeap.Enqueue(top)
	} else if this.minHeap.Size()-this.maxHeap.Size() == 2 {
		top, _ := this.minHeap.Dequeue()
		this.maxHeap.Enqueue(top)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := this.maxHeap.Size() - this.minHeap.Size()
	maxTop, _ := this.maxHeap.Peek()
	minTop, _ := this.minHeap.Peek()
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
