package leet_276

/*
There is a fence with n posts, each post can be painted with one of the k colors.

You have to paint all the posts such that no more than two adjacent fence posts have the same color.

Return the total number of ways you can paint the fence.

Note:
n and k are non-negative integers.

Example:

Input: n = 3, k = 2
Output: 6
Explanation: Take c1 as color 1, c2 as color 2. All possible ways are:

            post1  post2  post3
 -----      -----  -----  -----
   1         c1     c1     c2
   2         c1     c2     c1
   3         c1     c2     c2
   4         c2     c1     c1
   5         c2     c1     c2
   6         c2     c2     c1
*/

//如果第i个post和i-1染同一个颜色，那么第i-1个post不能和第i-2个post染同一个颜色；如果第i个post和第i-1个post染不同的颜色，那么第i个post的
//染色方案就是sum(i-1)*(k-1)。
//也就是为了得到i-th post的染色方案，我么需要利用之前的两个状态，一个是(i-1)-th 和(i-2)-th染一样颜色的方案数，另一个是(i-1)-th和(i-2)-th
//染不同颜色的方案数。对于这样的双重DP问题，我们可以使用两个数组来分别记录对应的状态，dp1表示和上一个post染同一个颜色的方案，dp2表示和上一个
//post染不同颜色的方案。递推公式如下：
//- dp1[i] = dp2[i-1]
//	表示如果i-th 和(i-1)-th染一样的颜色，那么染色方案为dp2[i-1]
//- dp2[i] = (dp1[i-1]+dp2[i-1])*(k-1)
//	表示i和i-1染不一样的颜色，对于(i-1)-th的每一种染色方案，我们有k-1种颜色可选

func numWays(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}

	dp1, dp2 := make([]int, n+1), make([]int, n+1)
	dp1[1], dp2[1] = 0, k
	for i := 2; i <= n; i++ {
		dp1[i] = dp2[i-1]
		dp2[i] = (dp1[i-1] + dp2[i-1]) * (k - 1)
	}

	return dp1[n] + dp2[n]
}

// 排列组合的方法，相邻的同颜色的post数量不超过2，那么要么全都不相同，要么存在几个相邻的post是同一颜色的。前者好计算，后者就比较费劲。
// 根据题意，当前post的颜色要么与上一个post的颜色不同，要么与上上个post的颜色不同。当前post的染色，如果与上一个post染同一个颜色，那么取决于
//上一个post有多少种和上上个post不同的染色方案；如果与上一个post染不同的颜色，那么取决于当前post和上一个post有多少个不同的染色方案，也就是
//（k-1)*上一个post的总染色方案。
func numWays2(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}

	same, diff, total := 0, k, k
	for i := 2; i <= n; i++ {
		same = diff            // ith post same as (i-1)th
		diff = total * (k - 1) // ith differ from (i-1)th
		total = same + diff
	}

	return total
}
