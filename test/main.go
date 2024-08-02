package main

import "fmt"

func minBombsToClear(nums []int, d int) int {
	bombs := 0
	for i := 0; i < len(nums); {
		if nums[i] > 0 {
			bombs += (nums[i]-1)/d + 1
			i += (nums[i] - 1) / d * d
		}
		i++
	}

	return bombs
}

func main() {
	nums := []int{5, 3, 2, 6, 8, 3}
	d := 2
	fmt.Println(minBombsToClear(nums, d))

}
