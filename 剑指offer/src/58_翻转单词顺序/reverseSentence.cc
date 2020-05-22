// 题目描述：输入一个英文句子，翻转句中单词的顺序，但单词内字符的顺序不变。例如“I am” 翻转为"am I"。
// 题解：
//   思路一：整体reverse，再逐个单词reverse。
//   思路二：用栈，整体入栈再出栈。然后逐个单词入栈再出栈。

#include <string>


class Solution {
public:
  void reverse(std::string& s, int begin, int end) {
    

    while(begin < end) {
      char temp = s[begin];
      s[begin] = s[end];
      s[end] = temp;

      begin++, end--;
    }
  }
    std::string ReverseSentence(std::string str) {
      if (str.empty()) return "";

      int begin = 0, end = str.length()-1;
      reverse(str, begin, end);
      begin = end = 0;

      while(begin != str.size()) {
        if (str[begin] == ' ') { // 跳过空格
          begin++;
          end++;
        }else if (str[end] == ' ' || end == str.length()) {
          reverse(str,begin, --end);
          begin = ++end;
        } else {
          end++;
        }
      }
      return str;
     }
};




// char* ReverseSentence(char* data) {
//   if (data == nullptr) return nullptr;

//   char* begin = data;
//   char* end = data;

//   while (*end != '\0') end++;
//   end--;

//   reverse(begin, end);

//   begin = end = data;

//   while(*begin != '\0') {
//           if (*begin == ' ') {
//             begin++;
//             end++;
//           }else if (*end == ' ' || *end == '\0') {
//             reverse(begin, end);
//             begin = ++end;
//           }else {
//             end++;
//           }
//   }
//   return data;
// }