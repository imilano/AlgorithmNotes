/*
题目描述：按照从小到大的顺序，求出第1500个丑数。丑数是只包含因子2、3、5的数。习惯上我们把1看做第一个丑数。


*/
#include <algorithm>

class Solution {
public:
 int GetUglyNumber_Solution(int index) { return getUglyNumber_v2(index); }

 bool isUgly(int number) {
   // 丑数就是能被2、3、5整除的数
   while (number % 2 == 0) number /= 2;
   while (number % 3 == 0) number /= 3;
   while (number % 5 == 0) number /= 5;

   return (number == 1) ? true : false;
}

// v1 暴力枚举，使用了大量无关的循环计算非丑数
int getUglyNumber_v1(int index) {
  int number=0;
  int uglyFound = 0;

  while(uglyFound < index) {
    if (isUgly(number)) uglyFound++;
    number++;
  }
  return number;
  }


// v2 创建数组保存已经找到的丑数，空间换时间，从而只计算丑数
//   创建一个数组，数组里保存的是已经排好序的丑数。如果让丑数排序呢？根据定义，丑数是一个丑数和2、3、5的乘积。
//   如果确保丑数是有序的呢？
//         1. 首先设一个已经排好序的丑数序列中的最大数是M
//         2. 接下来把数组的每个丑数乘以2。乘积要么是小于等于M的数，可直接舍去，要么是大于M的数，保存大于M中的最小数，记为M2，
//         3. 同理取得M3、M5。下一个数就是M2、M3、M5中的最小数。
//         4. 重复以上步骤，从而不断得到一个序列。
//         5. 2可进行优化。无需对每个丑数乘，因为总会有一部分乘积在M之内，所以只需要对特定的丑数序列做乘积即可
  int getUglyNumber_v2(int index) {
    if (index <= 0) return 0;
    int uglyNumber[index];
    uglyNumber[0] = 1;
    int nextUglyIndex = 1;

    int *pMulty2 =uglyNumber;
    int *pMulty3 = uglyNumber;
    int *pMulty5 = uglyNumber;


    while (nextUglyIndex < index) {
      int min = std::min(std::min(*pMulty2 * 2, *pMulty3 * 3), *pMulty5 * 5);
      uglyNumber[nextUglyIndex] = min;

      // 指针移动特定的距离，从而减少计算量
      while (*pMulty2 * 2 == uglyNumber[nextUglyIndex]) ++pMulty2;
      while (*pMulty3 * 3 == uglyNumber[nextUglyIndex]) ++pMulty3;
      while (*pMulty5 * 5 == uglyNumber[nextUglyIndex]) ++pMulty5;

      ++nextUglyIndex;
    }
  return uglyNumber[nextUglyIndex - 1];
  }
};
