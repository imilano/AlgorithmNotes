// 题目描述：统计一个数字在排序数组中出现的次数。
// 解题思路：使用二分查找算法，并且考虑到数组中会出现多个待查找数字的情况。只需要重复找到重复出现的数字的第一个k和最后一个k的位置即可。

#include <vector>
#include <algorithm>

class Solution {
public:
    int getFirstOfK(std::vector<int> data, int start, int end,int k) { // 找到左边第一个元素k的下标 
      if (start > end) return -1; // 不存在这样的k
      int mid = (start + end) / 2;
      int midData = data[mid];

      if (midData == k) {
              if ((data[mid-1] != k &&mid >0) || mid ==0 ) {  // 确保得到的是第一个k，而不是多个k里面的其中一个
                return mid;
              }else
                end = mid - 1;  // 如果是多个k里面的其中一个，那么继续查找
      }

      else if (midData > k) end = mid - 1;
      else if (midData < k) start = mid + 1;
      return getFirstOfK(data, start, end, k);
    }

    int getLastOfK(std::vector<int> data, int start, int end,int k) { // 找到右边第一个元素k的下标
      if (start > end) return -1;
      int mid = (start + end) / 2;
      int midData = data[mid];

      if (midData == k) {
              if ((mid <data.size()-1 && data[mid+1] != k) || mid == data.size()-1)
                return mid;
                else
                  start = mid + 1;
      }else if (midData > k)
        end = mid - 1;
        else if (midData < k)
          start = mid + 1;

        return getLastOfK(data, start, end, k);
    }
        // 二分查找算法
    int getNumberOkK_v1(std::vector<int> data, int k) {
        int number = 0; 
        
        if (!data.empty()) {
          int first = getFirstOfK(data, 0, data.size() - 1, k);
          int last = getFirstOfK(data, 0, data.size() - 1, k);

          if (first > -1 && last > -1) number = last - first + 1;
        }
        return number;
    }


//标准库equal_range的二分查找算法
    int getNumberOfK_v2(std::vector<int> data, int k) {
      auto pairsN = std::equal_range(data.begin(), data.end(), k);
      return pairsN.second - pairsN.first;
    }
    int GetNumberOfK(std::vector<int> data, int k) {
      if (data.empty()) return 0;
      return getNumberOfK_v2(data, k);
    }
};