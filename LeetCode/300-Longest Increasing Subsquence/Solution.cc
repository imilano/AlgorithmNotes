#include <vector>
#include <algorithm>

using std::fill;
using std::vector;
using std::max;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
      if (nums.empty()) return 0;
      int dp[nums.size()];
      fill(dp,dp+nums.size(), 1);  //将dp数组所有元素初始化为1

      for (int i = 0; i < nums.size();i++) 
        for (int j = 0; j < i;j++)
          if (nums[i] > nums[j]) dp[i] = max(dp[i], dp[j] + 1); // 如果当前值大于前面子数组中的任意一个值，那么就更新当前值的dp

      auto len = std::max_element(dp, dp + nums.size());
      return *len;
    }
};