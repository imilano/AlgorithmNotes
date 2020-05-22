// 题目描述： 输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。

// 1. 对数组排序，然后输出前k个数。
// 2. 基于快排中的partition函数，如果从第k个卡是partition，那么k前面都是比k小的，k后面都是比k大的 
// 3. 适合于海量数据的方式：开辟一个k大小的容器；遍历序列，如果容器未满，则压入元素；如果容器满了，那么拿当前元素和容器中的最大元素比较，如果比它小，那么就替换它，如果比它大，则继续下一个遍历。
//    数据结构最大堆适合这样的问题。使用STL中的set也可以实现。


#include <vector>
#include <algorithm>
#include<iostream>
#include <set>
class Solution {
  public:
        // 方法1, 排序思想
        void getLeaseK_v1(std::vector<int> input,std::vector<int>& output, int k) {
          std::sort(input.begin(), input.end());
          for (int i = 0; i < k; i++) {
            output.push_back(input[i]);
          }
          
        }
        void getLeastK_v2(std::vector<int> input,std::vector<int>& output, int k) {
                // 使用partition函数实现，借鉴快排思想
                int index = partition(input, 0, input.size() - 1, k);
                for (int i = 0; i < index; i++) {
                  output.push_back(input[i]);
                }
              }

        void getLeaseK_v3(std::vector<int> input,std::vector<int>& output, int k) {
                // std::max_element
                std::vector<int>::iterator it = input.begin();
                for (; it !=input.end();it++) {
                  if (output.size() < k) output.push_back(*it);
                  else {
                    std::vector<int>::iterator index = std::max_element(output.begin(),output.end());
                    if (*index > *it) {
                      output.erase(index);
                      output.push_back(*it);
                    }
                  }
                }
        }


        int partition(std::vector<int>& input, int start, int end, int k) {
          // int index = RandomInRange(start,end);  // 在start和end之间随机选择一个数
          int index = k;
          std::swap(input[index], input[end]);

          int small = start - 1;
          for (index = start; index < end; ++index) {
                  if (input[index] < input[end]) {
                    ++small;
                    if (small != index) std::swap(input[index], input[small]);
                  }
          }
          ++small;
          std::swap(input[small], input[end]);
          return small;
        }

    std::vector<int> GetLeastNumbers_Solution(std::vector<int> input, int k) {
      std::vector<int> output;
      if (k > input.size() || input.empty()) return output;
      getLeaseK_v1(input, output,k);
      return output;
    }
};