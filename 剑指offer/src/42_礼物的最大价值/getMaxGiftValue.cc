// 题目描述：在一个m*n的棋盘每一格都放有一个礼物，每个礼物都有一定的价值（价值大于0）。你可以从棋盘的左上角开始拿格子里的礼物，并且每次向左或者向下移动一格，直到到达棋盘的右下角。给定一个棋盘及其上面的礼物，请计算你最多可以拿到多少价值的礼物/
#include <vector>


// 采用动态规划方法。f(i,j)表示到达[i,j]的格子所能拿到礼物的最大值，从而有f(i,j) = max(f(i-1,j),f(i,j-1)) + gift[i][j]。

void getMaxValue_Solution(std::vector<std::vector<int> > values) {

}

int getMaxValue_v1(std::vector<std::vector<int> > values) {
  if (values.empty()) return 0;
  int rows = values.size();
  int cols = values[0].size();

  int dp[rows][cols];  // 过于浪费空间
  for (int i = 0; i < rows; i++) std::fill(dp[i], dp[i] + cols,0);  

  for (int i = 0; i < rows; i++)
    for (int j = 0; j < cols; j++) {
      int left = 0;
      int up = 0;

      if (i > 0) up = dp[i - 1][j];
      if (j > 0) left = dp[i][j - 1];
      dp[i][j] = std::max(left, up) + values[i][j];
    }

  return dp[rows - 1][cols - 1];
}


// 由于我们每次计算只需要用到f(i,j)左边（f(i,j-1)）和上边f(i-1,j)的元素，因此i-1行及其以上列的元素都没有必要保存下来。所以我们可以使用一个一维数组代替二维数组，
// 其长度为棋盘的列数cols，计算f(i,j)时，前面j个数字是f(i,0),f(i,1),f(i,2)...f(i,j-1)；数组从下标开始至最后一个数字 ，分别为f(i-1,j)，f(i-1,j+1),f(i-1,j+2)...f(i-1,n-1)
int getMaxValue_v2(std::vector<std::vector<int> > values) {
  int cols = values[0].size();
  int rows = values.size();
  if (values.empty()) return 0;

  int dp[cols];
  std::fill(dp, dp + cols, 0);
  for (int i = 0; i < rows; i++) {
    for (int j = 0; i < cols; j++) {
      int left = 0;
      int up = 0;

      if (i > 0) up = dp[j];  // dp[j]对应f(i-1,j)
      if (j > 0) left = dp[j - 1]; // dp[j-1]对应f(i,j-1)

      dp[j] = std::max(left, up) + values[i][j];
    }
  }
  return dp[cols - 1];
}
