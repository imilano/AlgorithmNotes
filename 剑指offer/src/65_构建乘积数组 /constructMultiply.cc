// 题目描述
// 给定一个数组A[0,1,...,n-1],请构建一个数组B[0,1,...,n-1],其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。（注意：规定B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2];）

// 题解：
//  可以把B[i] = A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]看成两部分（可以把A[i]看成1），一部分是C[i] = A[0]*A[1]*...*A[i-1]，另一部分是D[i] = A[i+1]*A[i+2]*...*A[n-1];
//  观察C[i]和D[i]可以看到，可以把C[i]自顶向下求值，即C[i]=C[i-1]*A[i-1]，其中C[0]=1；可以把D[i]自底向上i求值，其中D[i] = D[i+1]*A[i+1],其中D[n-1]等于1；


#include <vector>

class Solution {
public:
    std::vector<int> multiply(const std::vector<int>& A) {
      std::vector<int> B(A.size(), 0);
      if (A.empty()) return B;

      int len = A.size();
      B[0] = 1;
      for (int i = 1; i < len;i++) {
        B[i] = B[i - 1] * A[i - 1];
      }

      double temp = 1; // 复用B，Temp充当D[i]
      for (int i = len - 2; i >= 0;i--) {
        temp = temp * A[i + 1];
        B[i] *= temp;
      }
      return B;
    }
};