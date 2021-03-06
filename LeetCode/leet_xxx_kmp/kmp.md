KMP可以通过三种方式来解决。一种是暴力匹配，一种是通过部分匹配表，另一种是通过有限状态机。此处简介有限状态机的方式。



假设有模式串pat和待匹配的字符串str，二者的长度分别为P和S。

对于暴力匹配而言，我们逐一对str进行比较，直到连续P个字符都都和pat一样，那么我们就在str中找到了模式串。但是，存在找不到模式串的情况。对于暴力解法而言，假设模式串`pat[i]`和字符串`str[j]`不匹配时，我们就得回到模式串的开头，然后和字符串`str`的下一个字符开始继续进行比较。从而时间复杂度是`O(PS)`，空间复杂度`O(1)`。

暴力算法最不高效的地方就在于它暴力枚举，而我们知道，对于像`ABCAB`这样的模式串，当我们匹配到最后一个`B`发现不匹配时，我们是没有必要重置到开头来重新进行比较的。显而易见，当`B`不匹配时，`ABCA`是匹配的，而模式串中存在公共的`AB`这个公共前后缀，也就是说，下一步我们不需要完全重置到开头进行比较，而只需要从第一个`B`开始进行比较就够了。通过这样一个智能的办法，我们就减少了回溯的步骤，从而提高了比较的效率。

> 前缀： 字符串中不含最后一个字符的子串。例如A、AB是ABC的前缀。
>
> 后缀：字符串中不包含第一个字符的子串。例如B、BC是ABC的后缀。

我们从状态机的角度来理解这个问题。首先我们需要知道的是，字符串匹配中对我们最有价值的就是这个公共前后缀。有了这个公共前后缀，我们就知道当某个字符匹配失败时，我们下一步需要转移到的字符位置。KMP中的部分匹配表说到底，用的还是公共前后缀。

![](/Users/randy/go/src/lightsinger.life/algorithm/leetcode/leet_xxx_kmp/exp1.jpg)

而公共前后缀，其实只跟模式串有关。也就是说，对于任意一个模式串pat，只要模式串是确定的，那么无论str怎么变，我们计算出来的状态转移数组的值都不会变。

我们使用动态规划的思路来解决这个问题。首先，我们需要明确是什么决定了我们的状态转移。可以看到，状态转移有两个部分促成，一个是当前所处的状态，另一个是字符串str中待匹配的字符。我们 定义一个二维数组`dp[i][j]=m`，表示当前处于`i`状态，遇到字符串`j	`时需要转移到转态`m`。

字符的转移一共只会出现两种状态，假设当前处于转态`i`，需要匹配字符串`j`，那么：

- 模式串的下一个待匹配字符也是`j`，那么匹配成功，继续进行下一个字符的匹配，此时`dp[i][j]=i+1`。
- 模式串的下一个待匹配字符不是`j`，那么匹配失败，需要进行状态转移，此时`dp[i][j]=dp[X][j]`。其中`X`表示的是与当前字符具有相同前缀的字符所在的转态，表示的是在这个状态下遇到字符`j`需要转移到的位置。

对于匹配失败的情况，可以通过一个例子来加深一下理解。

假设模式串pat为`ABCAB`，当前匹配到了最后一个`B`发现跟str的字符不匹配，此时就需要转态转移，转移到哪里去呢？可以看到pat当前已经匹配的模式串`ABCAB`有一个公共的前后缀`AB`，也就说我们只需要转移到状态2就可以了，此时的数组应该更新为`dp[i][j]=dp[2][s]`，表示在2状态下遇到s应该转移到的状态。并且，由于X总是小于j，所以，等号右边的值总是已知的。

至此，可以将KMP分为两个部分，一个部分是计算出dp数组，第二部分根据计算出的dp数组来计算待查找字符串str是否匹配。



时间复杂度`O(N)`，N为字符串str长度，空间复杂度`O(M)`，M为模式串长度。

