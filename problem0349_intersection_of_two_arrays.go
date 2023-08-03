// https://leetcode.cn/problems/intersection-of-two-arrays
// 两个数组的交集
package main

func intersection(nums1 []int, nums2 []int) (ans []int) {
	set := map[int]struct{}{}
	res := map[int]struct{}{}
	addToSet(nums1, set)
	addToSet(nums2, res)
	if len(set) > len(res) {
		set, res = res, set
	}
	for v := range set {
		if _, ok := res[v]; ok {
			ans = append(ans, v)
		}
	}
	return
}

func addToSet(num []int, set map[int]struct{}) {
	for _, v := range num {
		set[v] = struct{}{}
	}
}
