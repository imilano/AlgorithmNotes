// 题目描述：和为s的两个数字。输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S。如果有多对数字的和等于S，则输出任意一对即可。

// 题解：双指针法。首指针从前向后遍历，尾指针从后向前遍历，每次相加二者之和，如果和小于S，那么前指针++，否则，后指针++，直到两个指针的元素为S，如果两个指针相遇，则返回错误。


#include <vector>

class Solution {
public:
    std::vector<int> FindNumbersWithSum(std::vector<int> array,int sum) {
      std::vector<int> result;

      if (array.empty()) return result;
      int start = 0, end = array.size() - 1;

      while(start < end) {
        int temp = array[start] + array[end];
        if (temp == sum) {
          result.push_back(array[start]);
          result.push_back(array[end]);
          break;
        } else if (temp < sum)
          start++;
          else if (temp > sum)
            end--;
      }
      return result;
    }
};