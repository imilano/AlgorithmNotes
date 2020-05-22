// 题目描述
//     给你一根长度为n的绳子，请把绳子剪成整数长的m段（m、n都是整数，n>1并且m>1），每段绳子的长度记为k[0],k[1],...,k[m]。请问k[0]xk[1]x...xk[m]可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

// 题解：
//     解法1：动态规划解法
//             定义f(n)为把长度为n的绳子剪成若干段后各段长度乘积的最大值。由于从上往下会涉及很多重复计算，因而我们采用自下而上的方法。
//             当n=2时，f(2) = 1;当n=3时间，f(3) = 2; 
//             方法一采用动态规划，每次重复计算自底向上计算，详见代码。
    
//     解法2：数学分析
//             当n>=5时，尽可能多剪长度为3的绳子；当剩下的绳子长度为4时，将绳子减成两段长为2的绳子。数学证明如下：
//                 当n>=5时，有2(n-2)>n 和 3(n-3)>n.也就是说，当剩下的绳子长度大于等于5时，尽量多剪长度2和3的绳子段。另外，当n>=5时，3(n-3)>2(n-2)，因而要多剪长度为3的绳子段。
//                 前面的假设是n>=5，那么当n为4呢，在长度为4的绳子上剪一刀有两种可能，1*3和2*2，由于后者更大，也就说当n为4时没必要剪，只是题目要求至少剪一刀。

#include<cmath>


class Solution {
public:

    int cutRope_dp(int length) {
        if (length < 2) return 0;
        if (length = 2) return 1; // 长度为2时最大为1
        if (length = 3) return 2;  // 长度为3时最大为2

        int* products = new int[length+1];
        // 乘积因子
        products[0] = 0; 
        products[1] = 1;
        products[2] = 2;
        products[3] = 3;

        int max = 0;
        for (int i=4 ; i<= length; i++) {
            for (int j=1; j<= i/2;j++) { // 枚举过程，O(n^2)
                if (products[j]*products[i-j] > max)
                    max = products[j]*products[i-j];
                products[i] = max;
            }
        }

        delete[] products;
        return products[length];

    }

    int cutRope_math(int len) {
        if (len <2) return 0;
        if (len == 2) return 1;
        if (len == 3) return 2;

        int timesOf3 = len/3; // 可以切分为多少个3
        // 当绳子长度为4时，把它切为2×2
        if (len - timesOf3*3 == 1) timesOf3--;
        int timesOf2 = (len-timesOf3*3)/2;
        return (int)(pow(2,timesOf2))*(int)(pow(3,timesOf3));

    }
    int cutRope(int number) {
        return cutRope_math(number);     
    }
};