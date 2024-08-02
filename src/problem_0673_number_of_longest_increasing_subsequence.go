// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/description/
// 最长递增子序列的个数
package main

import (
	"github.com/emirpasic/gods/v2/maps/treemap"
	"sort"
)

func findNumberOfLIS(nums []int) int {
	var dp []*treemap.Map[int, int]
	for _, num := range nums {
		cnt, size := 1, sort.Search(len(dp), func(i int) bool {
			return dp[i].Keys()[0] >= num
		})
		if size != 0 {
			p := dp[size-1]
			cnt = p.Values()[0]
			if _, value, ok := p.Ceiling(num); ok {
				cnt -= value
			}
		}
		if size == len(dp) {
			dp = append(dp, treemap.New[int, int]())
			dp[size].Put(num, cnt)
		} else {
			dp[size].Put(num, dp[size].Values()[0]+cnt)
		}
	}
	return dp[len(dp)-1].Values()[0]
}
