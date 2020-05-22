/*  题目：输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。
 *
 *  解法1,找规律：
 *      从前往后逐渐累加，使用一个max保存最大值，遇到负值时，比较加上之后的值是否比当前max大，如果大，更新max；如果当前sum小于0,那么令max=0,然后从当前值处开始继续累加计算。
 * 
 *  解法2,动态规划：
 *      用函数f(i)表示以第i个数字结尾的子数组的最大和，那么题目即为求max[f(i)]。则有
 *          当i=0或者f(i-1) <= 0 时，f(i)=pData[i]；当f(i-1)>0且i!=0时，f(i) = f(i-1) + pData[i]。
 * 
 */



#include <vector>
#include <iostream>
#include <algorithm>



class Solution {
public:
        // 找规律方法
    int FindGreatestSumOfSubArray(std::vector<int> array) {
      if (array.empty()) return 0;
      int curSum = 0;
      int maxSum = -1;
      int start = 0, end = 0;  // with index supported
      for (int i = 0; i < array.size(); i++) {
        if (curSum < 0) {  // 如果当前和小于0,那么指定当前和为arr[i]
                curSum = array[i];
                start = end = i;
        } else {
          curSum += array[i]; // 如果当前和大于0,则当前和的值为当前和加上arr[i]
          end++;
        }

        if (curSum > maxSum) maxSum = curSum;
      }

      std::cout << "Max from " << start << " to " << end << std::endl;
      for (int i = start; i < end; i++)  {
        std::cout << array[i] << " ";
      }
      return maxSum;
    }

//     // 动态规划方法
//     假设f(i) 表示以第i个数字结尾的子数组的最大和，那么我们需要求解出max(f(i)),我们有如下递归公式： 
//      如果 i = 0或者f(i - 1) <= 0，那么f(i) = arr[i];
//     如果i != 0 并且f(i - 1) > 0，那么f(i) = f(i - 1) + arr[i];

    int getMaxSumInSequence(std::vector<int> array) {
        int len = array.size();
        int maxSum[len];  // 记录动态规划中间值
        std::fill(maxSum, maxSum + len, 0);
        maxSum[0] = array[0];

        for (int i = 1; i < len;i++) {
          maxSum[i] = std::max(maxSum[i - 1] + array[i], array[i]);
        }

        auto s = std::max_element(maxSum, maxSum + len);
        return *s;
    }
};
