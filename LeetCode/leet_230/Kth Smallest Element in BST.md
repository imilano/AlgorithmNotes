### 题目描述
```
230. Kth Smallest Element in a BST
Medium

Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
```

### 题解

对于一个BST，它是天然有序的：对于每个节点，都有节点的值大于以该节点为根的左节点的值并且大于以该节点为根的右节点的值。对BST采用中序遍历，得出的序列必然是有序序列，结合栈和中序遍历，可以解答这个问题。