// 题目描述：
//         LL今天心情特别好,因为他去买了一副扑克牌,发现里面居然有2个大王,2个小王(一副牌原本是54张^_^)...他随机从中抽出了5张牌,想测测自己的手气,看看能不能抽到顺子,如果抽到的话,他决定去买体育彩票,嘿嘿！！“红心A,黑桃3,小王,大王,方片5”,“Oh My God!”不是顺子.....LL不高兴了,他想了想,决定大\小 王可以看成任何数字,并且A看作1,J为11,Q为12,K为13。上面的5张牌就可以变成“1,2,3,4,5”(大小王分别看作2和4),“So Lucky!”。LL决定去买体育彩票啦。 现在,要求你使用这幅牌模拟上面的过程,然后告诉我们LL的运气如何， 如果牌能组成顺子就输出true，否则就输出false。为了方便起见,你可以认为大小王是0。

//         要点：大小王可以看作任何牌，也就是说，可以把它插进不连续的牌中让牌连续。在数值上，可以把它看作0;

// 解题思路：先对数组进行排序，然后统计出0牌的数量num，最后统计排序之后的数组中相邻数字之间的空缺总数；如果空缺总数小于等于的个数，那么数字就是连续的，否则就是非连续的；
// 另外，还要注意，如果牌中出现了两张一模一样的牌，那么该牌就不可能是连续的。



#include <vector>
#include <algorithm>


class Solution {
public:
    bool IsContinuous(std::vector<int> numbers ) {
      if (numbers.empty()) return false;

      std::sort(numbers.begin(), numbers.end());

      int numOfZero = 0, numOfGap = 0;

      for (std::vector<int>::iterator it = numbers.begin();
           it != numbers.end();it ++) {
        if (*it == 0) numOfZero++;
      }

      int small = numOfZero; // small是排序后非0数字的下标
      int big = small + 1;   // 统计small到small+1之间的间隔; 然后将small和big不断向前移动
      while (big < numbers.size()) {
        if (numbers[small] == numbers[big]) return false;
        numOfGap += numbers[big] - numbers[small] - 1;
        small = big;
        big++;
      }

      return numOfZero >= numOfGap ? true : false;
    }
};