// 题目描述：左旋转字符串。
// 汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！

// 题解：三次reverse。前k个reverse一次，k+1到最后reverse一次，最后整个字符reverse一次。

#include <string>
#include <iostream>


class Solution {
public:
 void reverse(std::string& s, int start, int end) {
   char c; 
  while(start < end) {
    c = s[start];
    s[start] = s[end];
    s[end] = c;
    start++;
    end--;
  }
}
 std::string LeftRotateString(std::string str, int n) {
   if (str.empty() || n < 0 || n > str.length()) return "";

   int end = str.length() - 1;

   reverse(str, 0, n-1);
   reverse(str, n , end);
   reverse(str, 0, end);

   return str;
 }
};

// int main() { 
//   Solution s;
//   std::string t;
//   t = s.LeftRotateString("", 0);
//   return 0;
// }