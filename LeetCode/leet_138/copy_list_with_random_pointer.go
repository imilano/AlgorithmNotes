package leet_138

/*

A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a deep copy of the list.

The Linked List is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) where random pointer points to, or null if it does not point to any node.
*/


// Definition for a Node.
type Node struct {
    Val int
    Next *Node
    Random *Node
}

// Cause random could point to nil, for every new node, we can't traverse the entire list to find its random node. So we
// use a map to reduce time。
func copyRandomList1(head *Node) *Node {
    if head == nil {
        return head
    }

    m := make(map[*Node]*Node)
    res := &Node{
        Val:    head.Val,
        Next:   nil,
        Random: nil,
    }
    cur,node := head,res
    m[cur] = node

    // copy next
    cur = cur.Next
    for cur != nil {
        tmp := &Node{
            Val:    cur.Val,
            Next:   nil,
            Random: nil,
        }
        node.Next = tmp
        node = node.Next
        m[cur] = node
        cur = cur.Next
    }

    // copy random
    cur,node = head,res
    for cur != nil {
        node.Random = m[cur.Random]
        node = node.Next
        cur = cur.Next
    }

    return res
}

// recursive way
func copyRandomList2(head *Node) *Node {
    m := make(map[*Node]*Node)
    return helper(head,&m)
}

func helper(head *Node,m *map[*Node]*Node) *Node {
    if head == nil {
        return nil
    }

    // If existed mapping, return directly
    if v,ok := (*m)[head]; ok {
        return v
    }

    // get dup
    node := &Node{
        Val:    head.Val,
        Next:   nil,
        Random: nil,
    }
    (*m)[head] = node
    node.Next = helper(head.Next,m)
    node.Random = helper(head.Random,m)
    return node
}

// 多次扫描。第一趟直接将copy节点复制到当前节点的后面；第二趟扫描处理random节点；第三趟扫描将两个list分开
func copyRandomList(head *Node) *Node {
    if head == nil  {
        return head
    }

    // link the next
    cur := head
    for cur != nil {
        dup := &Node{
            Val:    cur.Val,
            Next:   cur.Next,
            Random: nil,
        }
        cur.Next = dup
        cur= dup.Next
    }

    // link the random
    cur = head
    for cur != nil {
        if cur.Random != nil {
            cur.Next.Random = cur.Random.Next
        }
        // WRONG
        //cur.Next.Random = cur.Random.Next
        cur = cur.Next.Next
    }

    // separate forked list from original list
    cur = head
    res := head.Next
    for cur != nil {  // careful
        t := cur.Next
        cur.Next = t.Next
        if t.Next != nil {
            t.Next = t.Next.Next
        }
        cur = cur.Next
    }

    return res
}
