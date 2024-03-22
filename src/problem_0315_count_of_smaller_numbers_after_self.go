// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/
// 计算右侧小于当前元素的个数
package main

type node struct {
	value int
	index int
}

func countSmaller(nums []int) []int {
	counts := make([]int, len(nums))
	arr := make([]node, len(nums))
	help := make([]node, len(nums))
	for i, num := range nums {
		arr[i] = node{num, i}
	}
	count(arr, help, 0, len(nums)-1, counts)
	return counts
}

func count(arr, help []node, l, r int, counts []int) {
	if l == r {
		return
	}
	m := (r + l) >> 1
	count(arr, help, l, m, counts)
	count(arr, help, m+1, r, counts)
	countMerge(arr, help, l, m, r, counts)
}

func countMerge(arr, help []node, l, m, r int, counts []int) {
	i, j, p := l, m+1, l
	for i <= m && j <= r {
		if arr[i].value <= arr[j].value {
			help[p] = arr[i]
			counts[arr[i].index] += j - m - 1
			i++
		} else {
			help[p] = arr[j]
			j++
		}
		p++
	}
	copy(help[p:r+1], arr[i:m+1])
	for i <= m {
		counts[arr[i].index] += j - m - 1
		i++
	}
	copy(help[p:r+1], arr[j:r+1])
	copy(arr[l:r+1], help[l:r+1])
}
