/* 题目描述
 *       请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"。
 *        当字符串中没有只出现一次的字符时，返回“#”。
 * 
 * 题解：
 *      使用一个哈希表，其key为字符，value为所出现的字符的下标。当一个字符再次出现时，更新其值为一个特殊值，例如-2。当一次扫描结束时，选择value中不等于-1且最小的那个value对应的key即为所求。
 */ 
#include <limits>

class Solution
{
  public:
    //Insert one char from stringstream
      void Insert(char ch)
      {
        if (occurrence[ch] == -1) occurrence[ch] = index;
        else if (occurrence[ch] >= 0)
          occurrence[ch] = -1;
      }
    //return the first appearence once char in current stringstream
      char FirstAppearingOnce()
      {
              char ch = '#';
              int minIndex = std::numeric_limits<int>::max();  // minIndex 指的是出现次数
              for (int i = 0; i < 256;i++) {
                      if (occurrence[i] >= 0 && occurrence[i] < minIndex) {  // 寻找出现次数最小的key
                        ch = (char)i;
                        minIndex = occurrence[i];
                      }
              }
              return ch;
      }
  Solution():index(0) {
        for (int i=0; i< 256; i++) {
        occurrence[i] = -1;
    }
  }

  private:
   int index;
   int occurrence[256];
};