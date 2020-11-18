## 题目描述

```
Imagine you have a special keyboard with the following keys:

Key 1: (A): Print one 'A' on screen.

Key 2: (Ctrl-A): Select the whole screen.

Key 3: (Ctrl-C): Copy selection to buffer.

Key 4: (Ctrl-V): Print buffer on screen appending it after what has already been printed.

Now, you can only press the keyboard for N times (with the above four keys), find out the maximum numbers of 'A' you can print on screen.

Example 1:

Input: N = 3
Output: 3
Explanation: 
We can at most get 3 A's on screen by pressing following key sequence:
A, A, A

 

Example 2:

Input: N = 7
Output: 9
Explanation: 
We can at most get 9 A's on screen by pressing following key sequence:
A, A, A, Ctrl A, Ctrl C, Ctrl V, Ctrl V

 

Note:

    1 <= N <= 50
    Answers will be in the range of 32-bit signed integer.

```

## 题解

采用递归的方式（或者备忘录方式）进行状态转移，很容易导致超时。

要得到最大的A数量，只有两种可能：
 - 要么是一直按A，此时一共有N次：A,A,...A（当 N 比较小时）
 - 要么是先按A，然后Ctrl A、Ctrl C、Ctrl V：A,A,...C-A,C-C,C-V,C-V,...C-V（当 N 比较大时）。
 
    这种情况下整个操作序列大致是：开头连按几个 A，然后 C-A C-C 组合再接若干 C-V，然后再 C-A C-C 接着若干 C-V，循环下去。
    
换句话说，最后一次按键要么是A要么是C-V，因而可以通过如下的方式来设计算法：
```
int[] dp = new int[N + 1];
// 定义：dp[i] 表示 i 次操作后最多能显示多少个 A
for (int i = 0; i <= N; i++) 
    dp[i] = max(
            这次按 A 键，
            这次按 C-V
        )
```

低于按A键的情况，`dp[i] = dp[i-1] + i`。而对于按C-V的情况，必然是CA、CC然后接着一串CV，我们使用j来指向若干CV的起点