// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 寻找两个正序数组的中位数
package main

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	size := n + m
	even := (size & 1) == 0
	if m != 0 && n != 0 {
		if even {
			return float64(findKthNum(nums1, nums2, size>>1)+findKthNum(nums1, nums2, (size>>1)+1)) / 2
		} else {
			return float64(findKthNum(nums1, nums2, (size>>1)+1))
		}
	} else if m != 0 {
		if even {
			return float64(nums1[(size-1)>>1]+nums1[size>>1]) / 2
		} else {
			return float64(nums1[size>>1])
		}
	} else if n != 0 {
		if even {
			return float64(nums2[(size-1)>>1]+nums2[size>>1]) / 2
		} else {
			return float64(nums2[size>>1])
		}
	} else {
		return 0
	}
}

func findKthNum(shorts, longs []int, k int) int {
	s, l := len(shorts), len(longs)
	if s > l {
		shorts, longs = longs, shorts
		s, l = l, s
	}
	if k <= s {
		return getUpMedian(shorts, 0, k-1, longs, 0, k-1)
	} else if k <= l {
		if longs[k-s-1] >= shorts[s-1] {
			return longs[k-s-1]
		}
		return getUpMedian(shorts, 0, s-1, longs, k-s, k-1)
	} else {
		if shorts[k-l-1] >= longs[l-1] {
			return shorts[k-l-1]
		}
		if longs[k-s-1] >= shorts[s-1] {
			return longs[k-s-1]
		}
		return getUpMedian(shorts, k-l, s-1, longs, k-s, l-1)
	}
}

func getUpMedian(a []int, s1, e1 int, b []int, s2, e2 int) int {
	for s1 < e1 {
		c1 := s1 + ((e1 - s1) >> 1)
		c2 := s2 + ((e2 - s2) >> 1)
		off := ((e1 - s1 + 1) & 1) ^ 1
		if a[c1] > b[c2] {
			e1 = c1
			s2 = c2 + off
		} else if a[c1] < b[c2] {
			s1 = c1 + off
			e2 = c2
		} else {
			return a[c1]
		}
	}
	return min(a[s1], b[s2])
}
