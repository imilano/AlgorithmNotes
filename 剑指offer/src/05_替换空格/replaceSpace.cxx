// 题目描述
// 请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

// 题解：先扫描一遍字符串，统计空格个数。每替换一个空格，字符串长度增加2.使用两个指针，一个指针P1指向原始字符的末尾，另一个指针P2指向替换之后字符的末尾。接下来向前移动P1,诸葛把它指向的字符复制到P2所指向的位置，直到碰到第一个空格为止，碰到第一个空格之后，P1向前移动一格，P2插入“%20”，并向前移动三格；重复上述过程，直到P1、P2指向同一个位置，表明所有空格都已经替换完毕。


#include <iostream>
#include <cstring>


using namespace std;


void replaceSpace(char *str,int length) {
                if ((str == nullptr)| length < 0)
                        return;
                int olength = 0, newLength = 0, numOfSpace = 0;
                char *p = str;

                // 统计空格个数
                while (*p != '\0')
                {
                        ++olength;
                        if (*p == ' ')
                                numOfSpace++;
                        p++;
                }

                // cout << "olength is " << olength
                //      << ", numOfSpace is " << numOfSpace << endl;
                // 完成替换之后的字符串长度
                newLength = olength + numOfSpace * 2;

                // 如果char数组空间不足，那么如何处理
                if (newLength > length)
                        return;
                // cout << "newLength is less than length" << endl;
                // 从后向前遍历
                int head = olength, tail = newLength;

                while(head >= 0 && tail > head) { // 终止条件，两个指针指向同一个位置
                        if (str[head] != ' ')
                        {
                                str[tail] = str[head];
                                tail--;
                        }
                        else
                        {
                                str[tail--] = '0';
                                str[tail--] = '2';
                                str[tail--] = '%';
                        }
                        head--;
                }
        }

// int main(){
//         char s[30] = " hello world, for you";
//         replaceSpace(s, 30);
//         cout << "string s is "<< s << endl;
//         return 0;
// }