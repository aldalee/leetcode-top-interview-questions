// https://leetcode.cn/problems/intersection-of-two-arrays
// 两个数组的交集
package main

func intersection(nums1, nums2 []int) []int {
	set := map[int]struct{}{}
	res := map[int]struct{}{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res[num] = struct{}{}
		}
	}
	return convert(res)
}

func convert(set map[int]struct{}) []int {
	ans := make([]int, 0, len(set))
	for num := range set {
		ans = append(ans, num)
	}
	return ans
}

// 面试题: 求两个有序数组的交集
func intersectionOfTwoOrderedArrays(a, b []int) (res []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			if res == nil || a[i] > res[len(res)-1] {
				res = append(res, a[i])
			}
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
