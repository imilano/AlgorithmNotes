// 题目描述：一只青蛙一次可以跳上1级台阶，也可以跳上2级台。求该青蛙跳上一个n级的台阶总共有多少种跳法。

// 题解：变相的斐波那契数列问题。
//         - 如果只有1级台阶，那显然只有1种跳法；
//         - 如果有2级台阶，那就有2种跳法：一种是分两次跳，另一种是一次跳两级
//         - 接下来考虑一般情况：
//                 - 把n级台阶跳法看成n的函数，记为f(n).
//                 - 由上述可知，f(1)=1,f(2)=2;
//                 - 当n>2时，最后一次跳的时候有两种不同的选择：
//                         - 一是最后一次只跳1级，此时跳法数量等于剩下的n-1级台阶的跳法数目，即为f(n-1);
//                         - 另一个是最后一次只跳2级，此时跳法数量等于剩下的n-2级台阶的跳法数目，即为f(n-2);
//                 - 综上所述，可得f(n) = f(n-1) + f(n-2)，可知为斐波那契数列。

class Solution {
public:
    int jumpFloor(int number) {
        if (number == 0)
                return 0;
        if (number ==  1)
                return 1;
        if (number == 2)
                return 2;
        int StepN = 0; // f(n)
        int StepNMinusOne = 2; // f(n-1)
        int stepNMinusTwo = 1; //f(n-2)

        for (unsigned int i = 3; i <= number;i++) {
                StepN = StepNMinusOne + stepNMinusTwo; // f(n) = f(n-1)+f(n-2)
                stepNMinusTwo = StepNMinusOne;  // f(n-2) = f(n-1),注意顺序和下一步不可颠倒
                StepNMinusOne = StepN; // f(n-1) = f(n)

        }
        return StepN;
    }
};