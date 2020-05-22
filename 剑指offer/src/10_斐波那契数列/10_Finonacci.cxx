// 题目描述：求斐波那契数列的第n项
//         写一个函数，输入n，求斐波那契数列的第n项。
// 
// 题解：基于递归的写法有太多的重复计算，因而使用自底向上的迭代。时间复杂度O(N)。

#include <iostream>
using namespace std;

class Solution {
public:
    int Fibonacci(int n) {
        if (n == 0)
            return 0;
        if (n == 1)
            return 1;
        int fibNMinusOne = 1;  //f(N-1)
        int fibNMinusTwo = 0;  //f(N-2)
        int fibN = 0; //f(N)
        for (unsigned int i = 2; i <= n;i++) {
            fibN = fibNMinusOne+fibNMinusTwo;
            fibNMinusTwo = fibNMinusOne;  // 注意次序
            fibNMinusOne = fibN;
            
        }
        return fibN;

    }
};

// int main() {
//         Solution s;
//         int f = s.Fibonacci(4);
//         cout << f << endl;
//         return 0;
// }

// class Solution {
// public:
//     int Fibonacci(int n) {
//         if(n == 0){
//             return 0;
//         }
//         if(n == 1){
//             return 1;
//         }
//         int a = 0,b = 1;
//         int m = 0;
//         int i;
//         for(i = 2;i <= n;i++){
//             m = a + b;
//             a = b;
//             b = m;
//         }
//         return m;
 
//     }
// };