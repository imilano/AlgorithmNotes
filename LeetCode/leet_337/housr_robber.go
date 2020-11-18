package leet_337


 // Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }


 var memo  = make(map[*TreeNode]int)

func max(a,b int) int {
    if a > b {
        return a
    }

    return b
}

func rob(root *TreeNode) int {
    if root == nil {
        return 0
    }

    // 备忘录消除重叠子问题
    if _,ok := memo[root]; ok {
        return memo[root]
    }

    // 抢，然后去下下家
    var do_rob,not_rob int
    do_rob = root.Val
    if root.Left == nil {
        do_rob += 0
    } else {
        do_rob += rob(root.Left.Left) + rob(root.Left.Right)
    }

    if root.Right == nil {
        do_rob += 0
    } else {
        do_rob += rob(root.Right.Right) + rob(root.Right.Left)
    }

    // 不抢，然后去下家
    not_rob = rob(root.Left) + rob(root.Right)

    res := max(do_rob,not_rob)
    memo[root] = res

    return res
}

 /*
  在递归解法中，之所以会有bad performance，是因为我们重复计算了子问题。而产生子问题的根本就在于我们rob函数的定义。对于一个根节点而言，它只会发生两种状态，一种是抢，一种是不抢。
  而我们的rob方法的定义完全丢失了这两个信息，因而导致了重复子问题的计算。因此，我们需要重新更改rob函数的签名，将其改为rob(root)[2]int，其返回两个值，第一个值表示不抢当前root节点所能得到的最大收益，
  第二个值表示抢当前root节点所能得到的最大收益。
  */

 func robNew(root *TreeNode) [2]int {
     if root == nil {
         return [2]int{0.0}
     }

     left := robNew(root.Left)
     right := robNew(root.Right)

     var res [2]int

     // 抢root，那么下家就不能抢了
     res[1] = root.Val + left[0] + right[0]

     // 不抢root，那么是否抢下家取决于下家给我带来的收益
     res[0] = max(left[0], left[1]) +
              max(right[0],right[1])

     return res
 }