// https://leetcode.cn/problems/merge-sorted-array
// 合并两个有序数组
package main

func Merge(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	for ; m > 0 && n > 0; idx-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m--
		} else {
			nums1[idx] = nums2[n-1]
			n--
		}
	}
	for ; m > 0; m-- {
		nums1[idx] = nums1[m-1]
		idx--
	}
	for ; n > 0; n-- {
		nums1[idx] = nums2[n-1]
		idx--
	}
}
