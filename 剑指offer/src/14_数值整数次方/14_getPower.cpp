// 题目描述
//         给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。保证base和exponent不同时为0。

// 题解：
//         999的二进制为1111100111,3^999 等于3^(512+256+128+644+63+4+2+1)，刚好与各二进制位对应。

#include <iostream>

using namespace std;

class Solution {
        public:
                /*
                 * 边界情况：
                 *      1. 指数为0;直接返回1
                 *      2. 指数为1; 直接返回base
                 *      3. 指数为负;先求指数为正的情况，然后将其作为分母
                 */
                double Power(double base, int exponent) {
                        if (exponent == 0)
                                return 1;
                        if (exponent == 1)
                                return base;
                        bool isNag = false;
                        if (exponent < 0) {
                                isNag = true;
                                exponent = -exponent;
                        }

                        // 快速幂求法
                        double result = 1;
                        while (exponent)
                        {
                                if (exponent&1)
                                        result *= base;
                                exponent >>= 1;
                                base *= base;
                        }

                        return isNag == true? 1/result:result;
                        // if (isNag) {
                        //         return 1 / result;
                        // } else
                        //         return result;
                }
};