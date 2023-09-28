// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
// 两个数组的交集II
package main

func intersect(nums1, nums2 []int) (ans []int) {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}
	return
}
