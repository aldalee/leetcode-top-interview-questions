// https://leetcode.cn/problems/number-of-islands/
// 岛屿数量
package main

func numIslands(grid [][]byte) (islands int) {
	row, col := len(grid), len(grid[0])
	set := make([]int, row*col)
	for i := range set {
		set[i] = i
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				idx := i*col + j
				islands++
				if j > 0 && grid[i][j-1] == '1' {
					a, b := find(set, idx), find(set, idx-1)
					if a != b {
						union(set, a, b)
						islands--
					}
				}
				if i > 0 && grid[i-1][j] == '1' {
					a, b := find(set, idx), find(set, idx-col)
					if a != b {
						union(set, a, b)
						islands--
					}
				}
			}
		}
	}
	return
}

func find(set []int, v int) int {
	if v != set[v] {
		set[v] = find(set, set[v])
	}
	return set[v]
}

func union(set []int, a, b int) {
	a, b = find(set, a), find(set, b)
	if a != b {
		set[a] = set[b]
	}
}
