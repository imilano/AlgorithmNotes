package leet_261

/*
Given n nodes labeled from 0 to n-1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

Example 1:

Input: n = 5, and edges = [[0,1], [0,2], [0,3], [1,4]]
Output: true
Example 2:

Input: n = 5, and edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
Output: false
Note: you can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0,1] is the same as [1,0] and thus will not appear together in edges.
*/

type pair struct {
	first, end int
}

//---------------------------------------
// 关键在于判断图联通且必须无环
// 根据题目所述，边不会冗余存储，故而如果这是有效的树的话，那么边的数量必然是节点的数量减1（如果数组中存储的都是有效节点的话）
func validTree(n int, edges []pair) bool {
	return len(edges) == n-1
}

// 并查集一直是解决连通性的有效方法，思想是遍历节点，如果两个节点相连，将其roots值连上，这样可以找到环。初始化roots数组为-1，然后对于一个pair的
//两个节点分别调用find函数，得到的值如果相同的话，说明环存在，返回false；不同的话，将其roots值union上。
func validTree2(n int, edges []pair) bool {
	roots := make([]int, n)
	for i := range roots {
		roots[i] = -1
	}

	for _, v := range edges {
		x, y := find(roots, v.first), find(roots,
			v.end) // 两个端点刚开始都不会出现在数组中，但是随着我们慢慢的加入和查找，一旦存在环，那么它的两个端点就会在并查集中重复出现，从而查找根时就会找到重复值。
		if x == y {
			return false
		}

		roots[x] = y
	}

	return len(edges) == n-1
}

func find(roots []int, i int) int {
	for roots[i] != -1 {
		i = roots[i]
	}

	return i
}

//-------------------------------
//DFS解法。使用DFS判断是否有环，使用一个数组visited判断是否联通，思路如下：
//首先根据边简历邻接链表，然后简历一个visited数组表示该节点是否已经被访问过。从节点0（或者其他任意节点）开始深度遍历，遍历过程如下：如果当前节
//点已经被访问过，那么说明重复访问，有环，直接返回false；如果没有被访问过，将当前节点标记为访问过，然后从邻接表中依次深度遍历当前节点的邻结点，
//继续递归遍历。需要注意的是，由于建立邻接表时，a是b的邻接节点，b是a的邻接节点，所以在遍历时，需要一个变量pre来记录上一个节点，以免回到上一个
//节点。遍历完成后，如果不返回false，说明图无环；然后查看visited，如果每个节点都被访问过，那么说明图联通。
func validTree3(n int, edges []pair) bool {
	visited := make([]bool, n)
	t := make([][]int, n)
	for _, v := range edges {
		t[v.first] = append(t[v.first], v.end)
		t[v.end] = append(t[v.end], v.first)
	}

	if !(dfs(0, -1, visited, t)) { // 如果当前节点有环
		return false
	}

	for k := range visited { // 如果当前有节点没有被访问过，说明图不连通
		if !visited[k] {
			return false
		}
	}

	return true
}

func dfs(cur int, pre int, v []bool, t [][]int) bool {
	if v[cur] {
		return false
	}

	v[cur] = true
	for _, n := range t[cur] {
		if n != pre { // 当前节点不跟上一个节点重合，避免发生重复访问
			if !dfs(n, cur, v, t) {
				return false
			}
		}
	}

	return true
}
