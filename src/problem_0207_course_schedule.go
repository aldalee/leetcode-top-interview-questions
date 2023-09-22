// https://leetcode.cn/problems/course-schedule/
// 课程表
package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		visited  = 0
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
		visited++
		for _, v := range edges[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return visited == numCourses
}
