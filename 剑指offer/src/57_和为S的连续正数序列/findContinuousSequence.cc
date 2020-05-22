// 题目描述
// 小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列? Good Luck!

// 题解：同上一题，使用两个指针，small指针指向1,
//     big指针指向2，对small到big之间的范围求和为curSum，如果curSum大于Sum，那么small++ ；如果curSum小于sum，big++;如果相等，那么就得到我们要找的值；


#include <vector>

class Solution {
public:
std::vector<std::vector<int> > result;

// 将small到big之间的数装入vector并返回
void getContinuousSequence(int start, int end) {  
  std::vector<int> r;
  for (int i = start; i <= end; i++) {
    r.push_back(i);
  }

  result.push_back(r);
}
std::vector<std::vector<int> > FindContinuousSequence(int sum) {
  int small = 1, big = 2;
  int middle = (sum + 1) / 2;

  if (sum < 3) return result;  // 正整数序列，并且至少需要包含两个数
  int curSum = small + big;

  while (small < middle) {
    if (curSum == sum) getContinuousSequence(small, big);
    while (curSum > sum && small < middle) {
      curSum -= small;
      small++;
      if (curSum == sum) getContinuousSequence(small, big);
    }
    big++;
    curSum += big;
  }
  return result;
    }
};