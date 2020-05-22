// 题目描述
// 每年六一儿童节,牛客都会准备一些小礼物去看望孤儿院的小朋友,今年亦是如此。HF作为牛客的资深元老,自然也准备了一些小游戏。其中,有个游戏是这样的:
// 首先,让小朋友们围成一个大圈。然后,他随机指定一个数m,让编号为0的小朋友开始报数。每次喊到m-1的那个小朋友要出列唱首歌,然后可以在礼品箱中任意的挑选礼物,
// 并且不再回到圈中,从他的下一个小朋友开始,继续0...m-1报数....这样下去....直到剩下最后一个小朋友,可以不用表演,并且拿到牛客名贵的“名侦探柯南”典藏版(名额有限哦!!^_^)。
// 请你试着想下,哪个小朋友会得到这份礼品呢？(注：小朋友的编号是从0到n-1).如果没有小朋友，请返回-1

// 题解：
//         方法一：使用一个环形链表来模拟圈。例如使用C++的list来模拟圈。由于list不是一个环形结构，所以在每次遍历到end时，我们需要自动把指针更正为begin。


#include <list>
class Solution {
public:

int getLastRaming_v1(int n, int m) {
  if (n < 1 || m < 1) return -1;
  unsigned int i = 0;
  std::list<int> numbers;
  for (i = 0; i < n; i++) numbers.push_back(i);

  std::list<int>::iterator current = numbers.begin();
  while (numbers.size() >1)
  {
    for (int j = 1; j < m; j++) {
      current++;
      if (current == numbers.end()) current = numbers.begin(); // 每次走m步
    }
    std::list<int>::iterator next = ++current;  
    if (next == numbers.end()) next = numbers.begin();
    --current;  // current已经自增，故而需要减1
    numbers.erase(current);
    current = next;
  }
  return numbers.front();
}

int LastRemaining_Solution(int n, int m) { return getLastRaming_v1(n, m); }
};