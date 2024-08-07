// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 数组中的第K个最大元素
package main

import (
	"math/rand/v2"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(arr []int, l, r, idx int) int {
	if l == r {
		return arr[l]
	}
	pivot := arr[l+rand.IntN(r-l+1)]
	L, R := partitions(arr, l, r, pivot)
	if idx < L {
		return quickSelect(arr, l, L-1, idx)
	} else if idx > R {
		return quickSelect(arr, R+1, r, idx)
	} else {
		return arr[idx]
	}
}

func partitions(arr []int, l, r, pivot int) (int, int) {
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

// TODO: bfprt
