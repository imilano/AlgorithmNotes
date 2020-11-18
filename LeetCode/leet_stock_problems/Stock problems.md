所有问题都基于以下这个问题：给定一个表示每天股票价格的数组，什么因素决定了可以获得的最大收益。

影响股票收益的因素有两个，第一个是时间，第二个是买卖次数。可以定义数组`dp[i][j]`，表示在剩余交易次数为`j`的情况下，当前第`i`天可以获得的最大利润。

另外，还需要注意的一点是，每次只能持有一支股票，也就是说，你当前要么持有股票，要么不持有股票，这也是一个状态。因而我们的数组可以更改为`dp[i][j][k]`，表示在还剩余`j`次买卖的情况下当前在第`i`天持有（k=1）或者不持有（k=0）所能拥有的最大利润。

基准条件如下：
```
dp[-1][k][0] = 0 // 还没有进行交易，所以利润为0

dp[i][0][0] = 0 // 不允许进行交易，所以利润为0

dp[-1][K][1] = -INF  // 未开始，不可能持有股票，用负无穷表示不可能

dp[i][0][1[ = -INF  // 不允许交易，也不可能持有股票，用负无穷表示不可能
```

需要注意的是，根据你当前是否持有股票的情况，也可以分别分出两中不同的状态来：

假设你当前持有股票，那么可能有以下两种方式：

- 你昨天也持有股票，今天没有购入股票。
- 你昨天不持有股票，但你今天购入了股票，所以你今天持有了股票。购入股票意味着你需要减去今天购入的值。

根据以上，当你持有股票时对应的状态转移方程是：`dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j-1][0]-price[i]`

同理，假设你当前不持有股票，那么也有两种状态：

- 你昨天也不持有股票，今天你也没有买入。
- 你昨天持有股票，但是你今天卖了。买了意味着你增加了你的利润，所有需要加上今天的股票值。

对应的状态转移方程是：`dp[i][j][0]=max(dp[i-1][j][0],dp[i-1][j][1] + price[i])`。

只有购入股票才会改变允许的最大交易次数，所以只有购入的时候才进行`j-1`。想求的最终答案是，拥有k次允许的交易次数，最后一天所能获取的最大利润。在最后一天，肯定是手里没有股票的收益大于手里有股票的收益。也就是说，最后求的结果是`result = dp[n-1][k][0]`。