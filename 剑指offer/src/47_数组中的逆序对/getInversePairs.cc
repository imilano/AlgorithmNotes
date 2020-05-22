// 题目描述：
//   在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。
//   并将P对1000000007取模的结果输出。 即输出P%1000000007
// 题解：
//    采用贵并排序的思想，先把数组分割为字数组，统计出 字数组内部的逆序对的数目，然后再统计出两个相邻子数组之间的逆序对的数目，在统计逆序对的过程中，
//    还需要对数组进行排序。举例来说，先把数组划分为单个元素，比较相邻两个元素之间的逆序对，然后合并两个元素并分别将其排序，成为一个长度为2的字数组，再比较两个字数组中逆序对的数目。
//    可以先使用两个指针分别指向两个数组的末尾，并每次只比较两个指针指向的数字。如果第一个字数组中的数字大于第二个，则逆序对的数量要加上第二个字数组的数量，并将指向第一个要字数组的指针向前移动一位，
//    如果第一个字数组中的数组小于等于第二个字数组中的数字，那么就不存在逆序对，此时将指向两个子数组的指针均向前移动一位。
//    
//    时间复杂度为O(nlogn)，因为需要一个数组存储排好序的数字，所以空间复杂度为战O(n)。
//
#include <vector>

class Solution {
public:

    int InversePairsCore(std::vector<int>& input,std::vector<int>& output,int start, int end) {
            if (start == end) {
              output[start] = input[start];
              return 0;
            }

            int length = (end - start) / 2;

            long long left = InversePairsCore(input, output, start, start + length);
            long long right = InversePairsCore(input, output, start + length+1,end);

            int i = start + length;
            int j = end;
            int copyIndex = end;
            long long count = 0;
            while (i >= start && j >= start + length + 1) {
              if (input[i] > input[j]) {
                output[copyIndex--] = input[i--];
                count += j - start - length;
              } else {
                output[copyIndex--] = input[j--];
              }
            }

            for (; i >= start; --i) output[copyIndex--] = input[i];
            for (; j >= start + length + 1; --j) output[copyIndex--] = input[j];

            return left + count + right;
    }
    int InversePairs(std::vector<int> data) {
      if (data.empty()) return 0;
      std::vector<int> copy(data);
//       for (int i = 0; i < data.size(); i++) copy.push_back(i);

      long long count = InversePairsCore(data, copy, 0, data.size() - 1);

      return count%1000000007;
    }
};