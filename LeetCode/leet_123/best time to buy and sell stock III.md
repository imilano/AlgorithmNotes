###  题目描述

```
123. Best Time to Buy and Sell Stock III


Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most two transactions.
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).


Example 1:

Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```

此处限制了k的值为2，从而每天有四个不同的转态变量。由于k的值比较小，我们可以把k的值枚举出来。然后再根据状态转移方程进行推导即可。

情况三和情况一相似，区别之处是，对于情况三，每天有四个未知变量：`T[i][1][0]、T[i][1][1]、T[i][2][0]、T[i][2][1]`，状态转移方程如下：
```go
T[i][2][0] = max(T[i - 1][2][0], T[i - 1][2][1] + prices[i])
T[i][2][1] = max(T[i - 1][2][1], T[i - 1][1][0] - prices[i])
T[i][1][0] = max(T[i - 1][1][0], T[i - 1][1][1] + prices[i])
T[i][1][1] = max(T[i - 1][1][1], T[i - 1][0][0] - prices[i]) = max(T[i - 1][1][1], -prices[i])
```
第四个状态转移方程利用了 T[i][0][0] = 0。
