// 题目描述：
//         输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。


#include <iostream>
using namespace std;

class Solution {
public:
     int  NumberOf1(int n) {
        //      int byteNumber = sizeof(n) * 8;

        //      if (n == 0)
        //              return 0;
        //      if (n < 0)
        //      {
        //              n = 1 << byteNumber + n;
        //      }
        //      int numberOfOne = 0;
        //      unsigned step = 1;
        //      while(byteNumber >= 0) {
        //              step = 1 << (byteNumber - 1);
        //              if ((n%step) == 1)
        //                      numberOfOne++;
        //              n = n << 1;
        //              byteNumber--;
        //      }

        //      return byteNumber;
        int count  =0;
        int flag =1;
        while (flag != 0 ){
                if ((n&flag) != 0)
                        count++;
                flag <<= 1;
        }
        return count;

        /*
         *  最优解:https://www.nowcoder.com/questionTerminal/8ee967e43c2c4ec193b040ea7fbb10b8?f=discussion
         * int count  = 0;
         * while (n != 0) {
         *      ++count;
         *       n=(n-1)&n;
         * }
         * return count;
         * 
         */
     }
};


int main(){
        Solution s;
        int t= s.NumberOf1(4);
        cout << t << endl;
        return 0;
}