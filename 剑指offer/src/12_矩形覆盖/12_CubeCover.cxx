// 题目描述：
//     我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？

// 题解：变相斐波那契数列问题。结合图形分析很容易得出结论。
//     相关链接：https://www.nowcoder.com/questionTerminal/72a5a919508a4251859fb2cfb987a0e6?f=discussion


class Solution {
public:
    int rectCover(int number) {
        if (number == 1)
            return 1;
        if  (number ==2 )
            return 2;
        int N = 0, NMinusOne = 2, NMinusTwo = 1;
        for (unsigned int i = 3; i <= number;i++){
                N = NMinusOne + NMinusTwo;
                NMinusTwo = NMinusOne;
                NMinusOne = N;
        }

        return N;
    }
};