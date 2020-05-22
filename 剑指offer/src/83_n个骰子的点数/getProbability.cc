/* 题目：把n个骰子仍在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
 * 题解：这是一道动态规划题。
 *      首先，扔n个骰子的解可以由n-1骰子的解推出，即原文提包含了子问题的解。其次，对于如果使用递归的话，那么必定村子重复的子问题，因此，本题适用于动态规划的解法。
 *      首先构造状态转移方程：设扔n个骰子时的排列数之和为s的函数为f(n，s)，那么如果先扔了n-1个骰子，则扔最后一个骰子就有6种情况（最后一个扔到1、2、3、4、5、6）来得到排列数之和s。
 *      即为f(n,s) = f(n-1,s-1)+f(n-1,s-2)+f(n-1,s-3)+f(n-1,s-4)+f(n-1,s-5)+f(n-1,s-6);当s>6n时，f(n,s)等于。初始状态位 n=1，f(1,1)=f(1,2)=f(1,3)=f(1,4)=f(1,5)=f(1,6)=1。
 *   
 *      详见：https://rongxuanhong.github.io/2018/09/26/%E5%89%91%E6%8C%87offer2%E4%B9%8Bn%E4%B8%AA%E9%AA%B0%E5%AD%90%E7%82%B9%E6%95%B0%E4%B9%8B%E5%92%8C/
 */ 
#include <iostream>
#include <cmath>

class Solution{
        public:
                int GetNumberCount(int num, int sum) {
                  if (num < 1 || sum > 6 * num || sum < num) return 0;
                  if (num == 1) return 1; // 只有一个骰子的时候，得到s的次数都只为那
                  
                  // 状态转移
                  int countsLeft = GetNumberCount(num - 1, sum - 1) +
                                   GetNumberCount(num - 1, sum - 2) +
                                   GetNumberCount(num - 1, sum - 3) +
                                   GetNumberCount(num - 1, sum - 4) +
                                   GetNumberCount(num - 1,sum - 5) +
                                   GetNumberCount(num - 1, sum - 6);
                  return countsLeft;
                }
               // 动态规划递归版
                void PrintProbability(int number) {
                  for (int i = number; i <= 6 * number;i++) { // 遍历可能的排列之和
                    std::cout << GetNumberCount(number, i) << std::endl;
                  }
                  // 计算每种排列之和对应的概率
                  int total = std::pow((float)6, number);
                  for (int i = number; i <= 6 *number;i++) {
                    float ratio = (float)GetNumberCount(number, i) / total;
                    std::cout << ":" << ratio << std::endl;
                  }
                }
};