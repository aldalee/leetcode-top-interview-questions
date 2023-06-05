package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for idx, num := range nums {
		if p, ok := m[target-num]; ok {
			return []int{p, idx}
		}
		m[num] = idx
	}
	return []int{-1, -1}
}
