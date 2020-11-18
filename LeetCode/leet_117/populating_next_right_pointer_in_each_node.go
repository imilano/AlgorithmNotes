package leet_117

/*
	Given a binary tree
	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}

	Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
	Initially, all next pointers are set to NULL.
	Follow up:
	    You may only use constant extra space.
	    Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

*/

//  Definition for a Node.
 type Node struct {
     Val int
     Left *Node
     Right *Node
     Next *Node
 }

 // Level traversal on every level, use two pointer, one points to a dummy head, another points to current visiting node.
func connect(root *Node) *Node {
    var result *Node
    dummyHead := new(Node)
    result = root
    for root != nil {
        cur := dummyHead
        for root != nil {
            if root.Left != nil {
                cur.Next = root.Left
                cur = cur.Next
            }
            if root.Right != nil {
                cur.Next = root.Right
                cur = cur.Next
            }
            root = root.Next
        }
        root = dummyHead.Next
        dummyHead = nil
    }

    return result
}