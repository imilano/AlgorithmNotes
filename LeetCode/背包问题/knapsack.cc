#include <algorithm>


const int N = 100;
int weight[N]={0};
int value[N] = {0};
int dp[N][N] = {0};  // 自动将所有元素初始化为0

int knapsack() {
        for ( int i =1;i <= N;i++)  // 物品数量
          for (int j = 1; j <= N;j++) // 容量数
          if (j < weight[i])  // 装不进
            dp[i][j] = dp[i - 1][j];
                else
                dp[i][j] =
                        std::max(dp[i-1][j], dp[i - 1][j - weight[i]] + value[i]);  // 装得进，取大者
        return dp[N][N];
}

int OptimizedKnapsack() {
        for (int i=1;i<= numbers;i++)
          for (int j = capacities; j >= 0;j--) 
          if j >= w[i] dp[j] = std::max(dp[j], dp[j - w[i]] + v[i]);

        return dp[N];
}