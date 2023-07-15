// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 数组中的第K个最大元素
// 要求时间复杂度O(n)
package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(arr []int, l, r, idx int) int {
	if l == r {
		return arr[l]
	}
	rand.Seed(time.Now().UnixNano())
	pivot := arr[l+rand.Intn(r-l+1)]
	L, R := partition(arr, l, r, pivot)
	if idx < L {
		return quickSelect(arr, l, L-1, idx)
	} else if idx > R {
		return quickSelect(arr, R+1, r, idx)
	} else {
		return arr[idx]
	}
}

func partition(arr []int, l, r, pivot int) (int, int) {
	less, more, cur := l-1, r+1, l
	for cur < more {
		if arr[cur] < pivot {
			less++
			arr[less], arr[cur] = arr[cur], arr[less]
			cur++
		} else if arr[cur] > pivot {
			more--
			arr[cur], arr[more] = arr[more], arr[cur]
		} else {
			cur++
		}
	}
	return less + 1, more - 1
}
