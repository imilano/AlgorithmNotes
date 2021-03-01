package leet_133

/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

For simplicity sake, each node's value is the same as the node's index (1-indexed). For example, the first node with val = 1, the second node with val = 2, and so on. The graph is represented in the test case using an adjacency list.

Adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.
*/

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	m := make(map[*Node]*Node)
	return dfs(&m, node)
}

// DFS
func dfs(m *map[*Node]*Node, root *Node) *Node {
	if root == nil {
		return nil
	}

	if _, ok := (*m)[root]; ok {
		return (*m)[root]
	}
	clone := new(Node)
	clone.Val = root.Val
	for _, node := range root.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(m, node))
	}
	(*m)[root] = clone
	return clone
}
