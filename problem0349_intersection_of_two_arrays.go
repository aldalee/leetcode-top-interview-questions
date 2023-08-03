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

// 面试题: 求两个有序数组的交集
func intersection2(a []int, b []int) (res []int) {
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
