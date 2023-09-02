// https://leetcode.cn/problems/subsets
// 子集
package main

func subsets(nums []int) (res [][]int) {
	var path []int
	var dfs func(idx int)

	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int(nil), path...))
		} else {
			dfs(idx + 1)
			path = append(path, nums[idx])
			dfs(idx + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return
}
