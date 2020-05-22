// 题目描述：数组中唯一只出现一次的数字。在一个数组中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现了一次的数字。

// 题解：还是用异或来进行计算。如果一个数字出现三次，那么它的二进制表示的每一位也出现三次。如果把所有出现三次的数字的二进制表示的每一位都分别加起来，那么每一位的和都能被3整除。
// 我们把数组中所有数字的二进制表示的每一位都加起来。如果某一位的和能被3整除，那么那个只出现一次的数字二进制对应的那一位是0,否则就是1。

#include <vector>
#include <iostream>
int findNumberAppearingOnce(std::vector<int> data) {
  if (data.empty()) return -1;

  int bitSum[32] = {0};
  for (int i = 0; i < data.size();i++) {
          // 逐个计算某一个数字的各个位是0还是1
    int bitMask = 1;
    for (int j = 31; j >= 0;--j) {
      int bit = data[i] & bitMask;
      if (bit != 0) bitSum[j] += 1;
      bitMask <<= 1;
    }
  }
  int result = 0;
  for (int i = 0; i < 32;i++) {
    result <<= 1;  // 0 左移多少位都是0
    result += bitSum[i] % 3;
  }
  return result;
}
