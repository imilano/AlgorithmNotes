### 题目描述
```
1143. Longest Common Subsequence
Medium

Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.

 

If there is no common subsequence, return 0.

 

Example 1:

Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

 

Constraints:

    1 <= text1.length <= 1000
    1 <= text2.length <= 1000
    The input strings consist of lowercase English characters only.
```

### 题解

定义一个二维数组`dp[i][j] = m`，表示对于字符串`s1[0...i]`和字符串`s2[0...j]`，二者的最长公共子串长度为`m`。那么有`base case`为`dp[0][i]`和`dp[j][0]`均为0，显然，空字符和非空字符的最长公共子串长度必然为0。
![](./dp.png)

动态规划中最难的就是状态转移，那么对于本题，状态转移方程应该是什么呢？假设`s1`和`s2`的最长公共子序列为`lcs`，那么对于`s1`和`s2`中的每个字符，有什么选择呢？显然，对于这样的字符，它要么在lcs中，要么不在lcs中。对于lcs中的每个字符，这个字符必然同时存在于`s1`和`s2`中。从而有：
```
使用双指针法从后向前遍历s1和s2，如果s1[i] == s2[j]，那么这个字符一定在lcs中，否则的话，s1[i]和s2[j]这两个字符至少有一个不在lcs中，需要先丢弃一个。
```
从而根据上述解法，可以得到如下的递归代码：
```
s1,s2 string
func dp(i,j) int {
    // 递归退出条件
    if i == -1 || j == -1 {
        return 0
    }

    // 如果相等，则继续计算子串
    if s1[i] == s2[j] {
        return dp[i-1][j-1] + 1
    } else {
        // 如果不相等，那么只需要取二者中的最大者即可
        return max(dp(i-1,j),dp(i,j-1))
    }
}

func lcs(s1,s2 string) int {
    return dp(len(s1)-1,len(s2)-1)
}
```
上述计算中，存在很多重复计算，只需要使用备忘录方法进行优化即可。