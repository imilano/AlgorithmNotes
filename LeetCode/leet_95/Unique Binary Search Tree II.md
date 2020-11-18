### 题目描述
```
95. Unique Binary Search Trees II   [Medium]

Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
```

本体采用与35题同样的解法，还是选取i作为根节点，同时小于i的数构成该根节点的左子树，大于i的节点构成该根节点的右子树。BST的要义就在于根节点大于左子树同时小于右子树，这杨的划分很自然的产生这样的结果。

我们定义 generateTrees(start, end) 函数表示当前值的集合为 [start,end][\textit{start},\textit{end}][start,end]，返回序列 [start,end][\textit{start},\textit{end}][start,end] 生成的所有可行的二叉搜索树。按照上文的思路，我们考虑枚举 [start,end][\textit{start},\textit{end}][start,end] 中的值 iii 为当前二叉搜索树的根，那么序列划分为了 [start,i−1][\textit{start},i-1][start,i−1] 和 [i+1,end][i+1,\textit{end}][i+1,end] 两部分。我们递归调用这两部分，即 generateTrees(start, i - 1) 和 generateTrees(i + 1, end)，获得所有可行的左子树和可行的右子树，那么最后一步我们只要从可行左子树集合中选一棵，再从可行右子树集合中选一棵拼接到根节点上，并将生成的二叉搜索树放入答案数组即可。

递归的入口即为 generateTrees(1, n)，出口为当 `start > end` 的时候，当前二叉搜索树为空，返回空节点即可。

