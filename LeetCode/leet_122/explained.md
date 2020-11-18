### 题目描述

```
122. Best Time to Buy and Sell Stock II

Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

```


与121题的不同之处在于，本题不限制交易的次数，k近似于无穷。当k近似于无穷时，k-1近似于k。则根据之前所述的状态转移方程，每天仍然有两个状态：

- 手头有股票： `dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])`
- 手头无股票： `dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - prices[i])`

由于k趋近于无穷，所有手头有股票的情况又可以精简为`dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k][0])`

可以看到，在手头有股票和手头无股票的情况下，`k`成为了一个常量，不在影响状态转移，所以状态转移方程中可以把k这一个维度省略。