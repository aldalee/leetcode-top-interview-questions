// Package shuffle_an_array
// https://leetcode.cn/problems/shuffle-an-array/
// 打乱数组
package shuffle_an_array

import "math/rand/v2"

type Solution struct {
	nums   []int
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:   nums,
		origin: append([]int(nil), nums...),
	}
}

func (s *Solution) Reset() []int {
	return s.origin
}

func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	for i := range s.nums {
		j := rand.IntN(n-i) + i
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
