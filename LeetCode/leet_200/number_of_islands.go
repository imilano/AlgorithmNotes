package leet_200

/*
Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]

Output: 1
*/

// DFS，仔细分析，类似于求图中的连通分量的个数。
// 使用一个visited数组来表示该位置是否被访问过，如果该点未被访问过并且为1，那么访问该点，然后继续访问与它相邻的1节点，直到找完相邻区域后，
// res自增1，然后继续寻找下一个未被访问过且数值为1的位置。
func numIslands(grid [][]byte) int {
	var res int
	if grid == nil || grid[0] == nil {
		return res
	}

	row,col := len(grid),len(grid[0])
	visited := make([][]bool,row)
	for i := range visited {
		visited[i] = make([]bool,col)
	}

	for i := 0; i< row;i++ {
		for j := 0; j<col;j++ {
			if grid[i][j] == '0' || visited[i][j] {
				continue
			}
			helper(grid,visited,i,j)
			res++
		}
	}

	return res
}

func helper(g [][]byte, v [][]bool,i,j int) {
	// 注意，row >= len(grid)就停止遍历，而不是row > len(grid)。此外，如果gid[row][col]为0也直接返回。
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[0]) || v[i][j] || g[i][j] == '0' {
		return
	}

	v[i][j] = true
	helper(g,v,i+1,j)
	helper(g,v,i-1,j)
	helper(g,v,i,j+1)
	helper(g,v,i,j-1)
}


// BFS。BFS与DFS真就孪生兄弟，形影不离。
func numIslandsBFS(grid [][]byte) int {
	var res int
	if grid == nil || grid[0] == nil {
		return res
	}

	row,col := len(grid),len(grid[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool,col)
	}


	dirX,dirY := []int{0,0,1,-1},[]int{1,-1,0,0}
	for i := 0; i<row;i++ {
		for j := 0; j < col;j++ {
			if visited[i][j] || grid[i][j] == '0' {
				continue
			}
			//visited[i][j] = true
			res++

			// 注意，是col，不是row
			que := []int{i*col+j}
			//que := []int{i*row+j}
			for len(que) != 0 {
				num := que[0]
				que = que[1:]
				for k := 0; k<4;k++ {
					x,y := num/col + dirX[k],num%col + dirY[k]
					if x < 0 || x >= row || y < 0 || y >= col || visited[x][y] || grid[x][y] == '0' {
						continue
					}
					visited[x][y] = true
					// que = append(que,x*row+y
					que = append(que,x*col+y)
				}
			}
		}
	}

	return res
}