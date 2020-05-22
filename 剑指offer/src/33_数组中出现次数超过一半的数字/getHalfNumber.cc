/*  题目描述
 *     数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。
 *
 *  题解：
 *    1. 基于数学上中位数的思想，对于一个排序的数组而言，数组中出现次数超过一半的数字必然为其中位数，但是数组的中位数并不一定就是出现一半的数字，因而需要再次对time进行检查。
 *    2. 基于数组本身的特征：数组中一个数字出现的次数超过数组长度的一半，就是说他出现的次数比其它数字加起来的次数都要多。
 *        因此，可以在遍历数组的时候保存两个值，一个是数组中的一个数字，一个是次数。如果我们们遍历到的下一个数字和之前保存的数字相同，那么次数加1；
 *        如果不同，那么次数减1；如果次数为0,那么我们保存该数字，并把次数设为1.
 * 
 */


#include <vector>
#include<algorithm>
class Solution {
public:
    int MoreThanHalfNum_Solution(std::vector<int> v) {
      if (v.empty()) return 0; // 如果不存在，则输出0

    }

    // 版本1 中位数思想
    int moreThanHalf_v1(std::vector<int> v) {
        // 确保有序
      sort(v.begin(), v.end());

      // 取中位数，有中位数并不能保证一定是我们所求的，因此需要检查该中位数出现过的次数
      int mid = v[v.size() / 2];
      int times = checkTimes(v, mid);
      if (times > v.size() / 2) return mid;
      else
        return 0;
    }

    // 版本2 数组本身特征
    int moreThanHalf_v2(std::vector<int> v) { 
        int times = 1;
        int temp = v[0];

        for (int i = 0; i < v.size; i++)  {
                if (times == 0) {
                  temp = v[i];
                  times = 1;
                }else if (v[i] == temp) {
                  times++;
                } else {
                  times--;
                }
        }
        int checkedTime = checkTimes(v, temp); // 这样的检查是必须的，以防数组中没有出现次数超过一半的数字
        if (checkedTime > v.size() / 2) return temp;
        else
          return 0;
    }

    int checkTimes(std::vector<int> v, int number) { 
            int time = 0;
      for (int i = 0; i < v.size();i++) {
        if (v[i] == number) time++;
      }
      return time;
    }

    // 
};