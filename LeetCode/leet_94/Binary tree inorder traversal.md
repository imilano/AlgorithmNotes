### 题目描述
```
94. Binary Tree Inorder Traversal
Medium

Given the root of a binary tree, return the inorder traversal of its nodes' values.

 

Example 1:

Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:

Input: root = []
Output: []

Example 3:

Input: root = [1]
Output: [1]

Example 4:

Input: root = [1,2]
Output: [2,1]

Example 5:

Input: root = [1,null,2]
Output: [1,2]

 

Constraints:

    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100

 

Follow up:

Recursive solution is trivial, could you do it iteratively?
```

方法1使用最常见的递归方式，详见代码，太过简单就不进一步叙述了。

方法二使用了栈，具体思想是，从根节点不断往左走，每遍历到一个节点就把该节点入栈，一直到树的最左节点。然后将栈顶节点弹出栈，访问该节点的值，之后继续遍历该节点的右子节点，详见代码。