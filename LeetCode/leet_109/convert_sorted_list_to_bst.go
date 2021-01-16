package leet_109

/*

Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of
every node never differ by more than 1.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --------------------------------------------------------------------------------------------------
// recursive way
// 由于原链表有序，故而只需要将单链表的中间节点作为根节点，然后将左边的节点作为根节点的左子树，右边的节点作为根节点的右子树，每次重复这个过程就可以
// 构造出一棵BST。需要注意的是，由于单链表之间有指向关系，所以为了方便递归，我们需要把中间节点和他之前的节点断开，故而我们使用一个pre指针来指示这个节点。
// 由于是单链表而不是数组，没办法很快找到中间节点，所以使用快慢指针，快指针每次走慢指针的一倍路程，当快指针走到末尾时，慢指针就可以走到中间。
// Time complexity: O(nlogn)
// Space complexity: O(logn)  ，递归调用空间复杂度永远不可能是O(1)
func findMiddle(head *ListNode) *ListNode {
	var pre *ListNode
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开pre和middle之间的链接
	if pre != nil {
		pre.Next = nil
	}

	return slow
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid := findMiddle(head)
	root := &TreeNode{Val: mid.Val}

	// If there is only one element
	if head == mid {
		return root
	}

	// mid和之前的节点已经断开，所以我们只需要将头结点传进去，就可以计算左子树
	root.Left = sortedListToBST(head)
	// mid 节点的下一节点就是右子树，因此也是一个递归的过程
	root.Right = sortedListToBST(mid.Next)
	return root
}

//-----------------------------------------------------
// time-space tradeoff
// 对于上一个解法，显然每次都需要进行遍历来找到中间元素，这极大的降低了效率，增加了时间复杂度。我们可以用一个数组来来存储单链表中的元素，
// 从而降低时间复杂度
// Time complexity: O(n), 每次找mid的时间降为了O(1)，而我们一共需要找n次
// Space complexity: O(n)
func convert(val []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: val[mid]}

	// If there is only one element , then just return
	if left == right {
		return root
	}

	root.Left = convert(val, left, mid-1)
	root.Right = convert(val, mid+1, right)
	return root
}

func sortedListToBST2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var val []int
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}

	return convert(val, 0, len(val)-1)
}

// ----------------------------------------------------
// Inorder simulation
var tmp *ListNode

func sortedListToBST3(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var size int
	tmp = head
	for head != nil {
		size++
		head = head.Next
	}

	return convertToBST(0, size-1)
}

func convertToBST(start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := convertToBST(start, mid-1)
	node := &TreeNode{Val: tmp.Val}
	node.Left = left

	tmp = tmp.Next
	node.Right = convertToBST(mid+1, end)
	return node
}
