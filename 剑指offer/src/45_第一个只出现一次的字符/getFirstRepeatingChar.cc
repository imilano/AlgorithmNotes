// 在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.

// 题解：使用一个hash表，存储会出现的字符，key为字符，value为其出现的次数。第一次扫描，value++，第二次扫描，输出value为1的key的下标。时间复杂度O(n)，空间复杂度O(1)

#include <string>

class Solution {
public:
    int FirstNotRepeatingChar(std::string str) {
      if (str.empty()) return -1;
      const int tableSize = 256;
      unsigned int table[tableSize];
      std::fill(table, table + 256, 0);
      int i = 0;
      while(str[i] != '\0') {
        table[str[i]]++;
        i++;
      }

      i = 0;
      while(str[i] != '\0') {
        if (table[str[i]] == 1) {
          int j = str.find_first_of(str[i]);
          return j;
        }
        
        i++;
      }
      return -1;
    }
};