package leet_286

/*
You are given a m x n 2D grid initialized with these three possible values.

-1 - A wall or an obstacle.
0 - A gate.
INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.

For example, given the 2D grid:
INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
After running your function, the 2D grid should be:
  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
*/

// 规定-1是墙，0是门，求每个点到门的最近距离
// DFS。首先找到一个值为0的点，然后对该点的前后左右进行深度遍历。同时代入深度值为1，如果当前值大于深度值，那么更新当前值为深度值，并对当前点
//的四个点继续深度遍历，注意将深度值加1
func wallsAndGates(rooms [][]int) {
	row, col := len(rooms), len(rooms[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rooms[i][j] == 0 {
				dfs(rooms, i, j, 0)
			}
		}
	}
}

func dfs(rooms [][]int, i, j, val int) {
	if i < 0 || i >= len(rooms) || j < 0 || j >= len(rooms[0]) || rooms[i][j] < val {
		return // 注意，这里的rooms[i][j]<val其实已经包含了遇到墙（-1）的情况，因为val的值是从0开始增加的
	}

	rooms[i][j] = val
	dfs(rooms, i, j+1, val+1)
	dfs(rooms, i, j-1, val+1)
	dfs(rooms, i-1, j, val+1)
	dfs(rooms, i+1, j, val+1)
}

// BFS
type pair struct {
	i, j int
}

func wallsAndGates2(rooms [][]int) {
	var que []pair
	row, col := len(rooms), len(rooms[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rooms[i][j] == 0 {
				que = append(que, pair{
					i: i,
					j: j,
				})
			}
		}
	}

	var cords = []pair{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for len(que) != 0 {
		cur := que[0]
		que = que[1:]

		curI, curJ := cur.i, cur.j
		for i := 0; i < len(cords); i++ {
			nextI := curI + cords[i].i
			nextJ := curJ + cords[i].j

			if nextI < 0 || nextI >= row || nextJ < 0 || nextJ >= col || rooms[nextI][nextJ] < rooms[curI][curJ]+1 {
				continue
			}
			rooms[nextI][nextJ] = rooms[curI][curJ] + 1
			que = append(que, pair{nextI, nextJ}) // 为什么需要这一步？可以举例说明，假设一整个矩阵中只有中心点是gate，那么
			//我们如何更新中心点周围的所有点呢？就只能通过更新了周围点之后，将周围点作为新的中心点，然后继续遍历
		}
	}
}
