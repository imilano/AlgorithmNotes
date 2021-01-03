package leet_207

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should
also have finished course 1. So it is impossible.


Constraints:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
1 <= numCourses <= 10^5
*/

// BFS
// 环检测问题，使用邻接表建立起有向图，同时新建一个入度数组，采用深度遍历的方式，建立一个队列，将入度为0的节点放入队列中。
// 开始先根据输入来建立这个有向图，并将入度数组也初始化好。然后定义一个 queue 变量，将所有入度为0的点放入队列中，然后开始遍历队列，
// 从 graph 里遍历其连接的点，每到达一个新节点，将其入度减一，如果此时该点入度为0，则放入队列末尾。
// 直到遍历完队列中所有的值，若此时还有节点的入度不为0，则说明环存在，返回 false，反之则返回 true
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses) // 建立邻接表
	in := make([]int,numCourses)
	for _,v := range prerequisites {
		graph[v[1]] = append(graph[v[1]],v[0])
		in[v[0]]++
	}

	var que []int   // 如果一个节点的入度为0， 则该节点没有依赖，可以作为起始节点，将该节点入队
	for i := 0; i< numCourses;i++ {
		if in[i] ==  0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {  // 取出队中的起始节点
		top := que[0]
		que = que[1:]

		for _,v := range graph[top] {
			in[v]--
			if in[v] == 0 {
				que = append(que, v)
			}
		}
	}

	for i := 0; i< numCourses;i++ {
		if in[i] != 0 {
			return false
		}
	}

	return true
}