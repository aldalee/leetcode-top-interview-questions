// https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters
// 至多包含K个不同字符的最长子串
// Given a string,find the length of the longest substring T that contains at most k distinct characters.
// Example 1:
//
//	Input:s="eceba" k=2
//	Output:3
//	Explanation: T is "ece" which its length is 3.
//
// Example 2:
//
//	Input:s="aa" k=1
//	Output:2
//	Explanation: T is "aa" which its length is 2.
package main

func lengthOfLongestSubstringKDistinct(s string, k int) (ans int) {
	cnt := make([]int, 256)
	diff, n := 0, len(s)
	for l, r := 0, 0; l < n; l++ {
		for r < n && (diff < k || (diff == k && cnt[s[r]] > 0)) {
			if cnt[s[r]] == 0 {
				diff++
			}
			cnt[s[r]]++
			r++
		}
		ans = max(ans, r-l)
		if cnt[s[l]] == 1 {
			diff--
		}
		cnt[s[l]]--
	}
	return
}
