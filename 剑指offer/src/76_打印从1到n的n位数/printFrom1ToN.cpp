// 题目描述：输入数字n，按顺序打印出从1到最大的n位十进制数。比如输入3,则打印出从1、2、3一直到最大的3位数999

// 题解：
//     n位十进制数其实就是n个从0到9的全排列，也就是说 ，把数字的每一位都从0到9排列一遍，就得到了所有的十进制数。只是在打印的时候，排在前面的0u不打印出来罢了。用递归表达全排列，递归结束的条件是我们已经设置了数字的最后一位
#include <cstring>
#include<cstdio>
#include <cstdlib>

void printNumber(char* number) {
    bool isBegining0 = true;
    int nlength = strlen(number);

    for ( int i=0; i < nlength; i++) {
        if (isBegining0 && number[i] != '0') isBegining0 = false;

        if (!isBegining0) {
            printf("%c",number[i]);
        }
    }
    printf("\t");
}

void printToMaxOfNDigitsRecursively(char* number,int length, int index) {
    if (index = length-1) {  // 递归结束条件，已经设置到了数字的最后一位
        printNumber(number);
        return;
    }

    for ( int i=0;  i< 10; i++) {
        number[index+1] = i+'0';
        printToMaxOfNDigitsRecursively(number,length,index+1);
    }
}

void print1ToMaxOfDigits(int n) {
    if (n < 0) return ;

    char* number = new char[n+1];  // n+1 是因为最后是一个结束符'\0'
    number[n] = '\0';

    for ( int i =0; i< 10;i++) {
        number[0]=i + '0';
        printToMaxOfNDigitsRecursively(number,n,0);
    }

    delete[] number;
}