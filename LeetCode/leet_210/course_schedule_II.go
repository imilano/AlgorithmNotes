package leet_210

/*
There are a total of n courses you have to take labelled from 0 to n - 1.

Some courses may have prerequisites, for example, if prerequisites[i] = [ai, bi] this means you must take the course bi before the course ai.

Given the total number of courses numCourses and a list of the prerequisite pairs, return the ordering of courses you should take to finish all courses.

If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int

	graph := make([][]int, numCourses)
	in := make([]int, numCourses)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		in[v[0]]++
	}

	var que []int
	for i := range in {
		if in[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {
		top := que[0]
		que = que[1:]

		res = append(res, top)
		for _, v := range graph[top] {
			in[v]--
			if in[v] == 0 {
				que = append(que, v)
			}
		}
	}

	if len(res) != numCourses {
		res = nil
	}

	return res
}
