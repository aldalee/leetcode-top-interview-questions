// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/
// 计算右侧小于当前元素的个数
// TODO: 待弄懂，树状数组的解法
package main

import "slices"

func countSmaller(nums []int) []int {
	minVal := slices.Min(nums)
	maxVal := slices.Max(nums)
	n := maxVal - minVal + 1
	tree := make([]int, n<<1)
	cnts := make([]int, len(nums))
	add := func(i, v int) {
		for ; i <= n; i += i & -i {
			tree[i] += v
		}
	}
	sum := func(i, idx int) {
		ans := 0
		for ; i > 0; i -= i & -i {
			ans += tree[i]
		}
		cnts[idx] = ans
	}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i] - minVal + 1
		add(num, 1)
		sum(num-1, i)
	}
	return cnts
}
