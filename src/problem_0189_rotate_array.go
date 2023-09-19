// https://leetcode.cn/problems/rotate-array/
// 旋转数组
package main

// Rotate please modify to rotate before submitting.
func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	var reverse = func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
