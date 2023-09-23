// https://leetcode.cn/problems/contains-duplicate/
// 存在重复元素
package main

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
