// 题目描述：假设把某股票按照时间先后顺序存储在数组中，请问买卖该股票一次性可能获得的最大利润是多少？
// 例如，一只股票在某些时间节点的价格为{9,11,8,5,7,12,16,14}。如果我们能在价格为5的时候买入并在价格为16时候卖出，则收获最大的利润11

// 题解：最大的利润就是数组中所有数对的最大差值。定义函数diff为当卖出价为数组中第i个数字时可能获得的最大利润。
// 显然，当卖出价固定时，买入价越低获得的利润越大。也就是说，如果我们在扫描第i个数字时，能够记住之前的i-1个数字中的最小值，
//  就能算出在当前价位卖出时可能得到的最大利润。



#include <vector>

// 时间复杂度O（n）
int maxProfit(std::vector<int> data) {
  if (data.empty() || data.size() < 2) return -1;

  int min = data[0];
  int maxDiff = data[1] - min;

  for (int i = 2; i < data.size();++i) {
    if (data[i - 1] < min) min = data[i - 1];

    int curDiff = data[i] - min;
    if (curDiff > maxDiff) maxDiff = curDiff;
  }

  return maxDiff;
}