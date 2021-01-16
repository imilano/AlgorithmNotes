package leet_116

//  Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var pre, cur *Node
	pre = root
	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}

	return root
}

// Loop from level 0 to n-2, in every level, traverse this level and connect children
func connectSleet(root *Node) *Node {
	// for loop will lose reference to root, so we need a new node to store root node.
	result := root
	for root != nil && root.Left != nil {
		cur := root
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}

		root = root.Left
	}

	return result
}
