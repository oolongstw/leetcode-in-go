package main

import "fmt"

// 判断是否可能完成所有课程的学习
// 解题思路: 通过「拓扑排序」判断 Graph 是否是有向无环图(DAG)
// 关键点: 各个节点的入度 + 邻接表
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录每个节点的入度
	indegree := make([]int, numCourses)
	// 邻接表
	m := map[int][]int{}
	for _, item := range prerequisites {
		// item: [u, v] u->v u到v的一条有向边 v的入度+1, u的出度+1, 想学v必须要先学u
		u, v := item[0], item[1]
		indegree[v]++
		// 构造邻接表
		if _, ok := m[u]; !ok {
			m[u] = []int{v}
		} else {
			m[u] = append(m[u], v)
		}
	}
	// BFS queue
	queue := []int{}
	for idx, vert := range indegree {
		if vert == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) != 0 {
		head := queue[0]
		adj := m[head]
		queue = queue[1:]
		numCourses--
		for _, v := range adj {
			indegree[v]-- // 所有邻接节点入度-1
			if indegree[v] == 0 {
				queue = append(queue, v) // enqueue
			}
		}
	}

	// 因此，拓扑排序出队次数等于课程个数，返回 numCourses == 0 判断课程是否可以成功安排。
	// return numCourses == 0

	// 若是有向无环图 则所有节点入度最终都为0
	for _, v := range indegree {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 2
	prerequisites := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println(canFinish(num, prerequisites))
}
