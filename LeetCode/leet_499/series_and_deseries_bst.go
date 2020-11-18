package leet_499

/*
	Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer,
	or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
	Design an algorithm to serialize and deserialize a binary search tree. There is no restriction on how your serialization/deserialization
	algorithm should work. You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.

	The encoded string should be as compact as possible.
 */


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// TODO figure it out
type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    return nil
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
