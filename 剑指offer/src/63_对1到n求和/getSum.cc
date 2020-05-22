// 题解：求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

class Solution {
 public:
  Solution() {
    N++;
    Sum += N;
  }
  int Sum_Solution(int n) {
    Solution::reset();
    Solution* s = new Solution[n];
    return Solution::getSum();
  }
  static void reset() {
    N = 0;
    Sum = 0;
  }

  int getSum() { return Sum; }

 private:
  static unsigned int N;
  static unsigned int Sum;
};

unsigned int Solution::N = 0;
unsigned int Solution::Sum = 0;
