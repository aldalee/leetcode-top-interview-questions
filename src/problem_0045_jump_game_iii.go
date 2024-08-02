// Given a 0 indexed integer array nums of length n.
// The initial position is start, and the target position is end.
// and jumping is not allowed when it exceeds the limit.
// Please return the minimum number of jumps from start to end.
// If it is not reachable, return -1.
// Note: 0<=start<=n-1, 0<=end<=n-1

package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

// Solution: BFS (optimal solution)
func minJumpTimes(nums []int, n, start, end int) int {
	if start < 0 || start >= n || end < 0 || end >= n {
		return -1
	}
	queue := list.New()
	queue.PushBack(start)
	levelMap := make(map[int]int) // <index, level>
	levelMap[start] = 0

	for queue.Len() > 0 {
		idx := queue.Remove(queue.Front()).(int)
		level := levelMap[idx]
		if idx == end {
			return level
		} else {
			left := idx - nums[idx]
			right := idx + nums[idx]
			if left >= 0 && levelMap[left] == 0 {
				queue.PushBack(left)
				levelMap[left] = level + 1
			}
			if right < n && levelMap[right] == 0 {
				queue.PushBack(right)
				levelMap[right] = level + 1
			}
		}
	}
	return -1
}

// Solution: DP
func dp(nums []int, n int, start int, end int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -2
		}
	}
	return Jump(nums, n, start, end, 0, dp)
}

func Jump(nums []int, n, idx, end, step int, dp [][]int) int {
	if idx == end {
		dp[idx][step] = step
		return step
	}
	if idx < 0 || idx >= n || step > n-1 {
		return -1
	}
	if dp[idx][step] != -2 {
		return dp[idx][step]
	}
	p1 := Jump(nums, n, idx-nums[idx], end, step+1, dp)
	p2 := Jump(nums, n, idx+nums[idx], end, step+1, dp)
	ans := -1
	if p1 != -1 && p2 != -1 {
		ans = min(p1, p2)
	} else if p1 != -1 {
		ans = p1
	} else if p2 != -1 {
		ans = p2
	}
	dp[idx][step] = ans
	return ans
}

func generateRandomArray(N, V int) []int {
	n := rand.Intn(N) + 1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(V)
	}
	return arr
}

func main() {
	maxLen := 1000
	maxValue := 1000
	times := 10000
	ok := true
	for i := 0; i < times; i++ {
		nums := generateRandomArray(maxLen, maxValue)
		n := len(nums)
		start := rand.Intn(n)
		end := rand.Intn(n)
		expected := minJumpTimes(nums, n, start, end)
		output := dp(nums, n, start, end)
		if expected != output {
			fmt.Println("Wrong Answer!")
			fmt.Println("nums =", nums)
			fmt.Println("n =", n)
			fmt.Println("start =", start)
			fmt.Println("end =", end)
			fmt.Println("Expected =", expected)
			fmt.Println("Output =", output)
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("Accepted!")
	}
}
