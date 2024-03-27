// https://leetcode.cn/problems/wiggle-sort-ii/
// 摆动排序II
package main

import "math/rand/v2"

func wiggleSort(nums []int) {
	n := len(nums)
	partitionAroundPivot(nums, 0, n-1, n>>1)
	if n&1 == 0 {
		perfectShuffle(nums, 0, n-1)
		reverseArray(nums, 0, n-1)
	} else {
		perfectShuffle(nums, 1, n-1)
	}
}

func partitionAroundPivot(nums []int, l, r, idx int) {
	for l < r {
		pivot := nums[rand.IntN(r-l+1)+l]
		left, right := netherlandsFlag(nums, l, r, pivot)
		if left <= idx && idx <= right {
			break
		} else if idx < left {
			r = left - 1
		} else {
			l = right + 1
		}
	}
}

func netherlandsFlag(arr []int, l, r, pivot int) (int, int) {
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

func perfectShuffle(arr []int, l, r int) {
	for r-l+1 > 0 {
		shift, k := maxInShuffle(r - l + 1)
		mid := l + (r-l)>>1
		rotateArray(arr, l+shift>>1, mid, mid+shift>>1)
		cycleLeader(arr, l, shift, k)
		l += shift
	}
}

func maxInShuffle(remaining int) (int, int) {
	base, k := 3, 1
	for base <= (remaining+1)/3 {
		base *= 3
		k++
	}
	return base - 1, k
}

func rotateArray(arr []int, l, m, r int) {
	reverseArray(arr, l, m)
	reverseArray(arr, m+1, r)
	reverseArray(arr, l, r)
}

func reverseArray(arr []int, l, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func cycleLeader(arr []int, start, n, k int) {
	idx := 1
	for i := 0; i < k; i++ {
		shiftVal := arr[idx+start-1]
		next := nextIdx(idx, n)
		for next != idx {
			shiftVal, arr[next+start-1] = arr[next+start-1], shiftVal
			next = nextIdx(next, n)
		}
		arr[next+start-1] = shiftVal
		idx *= 3
	}
}

func nextIdx(i, n int) int {
	return (2 * i) % (n + 1)
}
