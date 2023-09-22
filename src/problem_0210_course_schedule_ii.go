// https://leetcode.cn/problems/course-schedule-ii/
// 课程表II
package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		order    []int
		queue    []int
	)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indegree[info[0]]++
	}
	for v := 0; v < numCourses; v++ {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range edges[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}
