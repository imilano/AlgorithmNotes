/*
题目描述：从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

题解：动态规划
        1.
定义f(i)表示以第i个字符为结尾的不包含重复字符的子字符的最长长度，因为我们已经知道f(i-1)了，所以我们只需要查询第i个字符是否出现在了f(i-1)中即可知是否重复。
        2. 如果第i个字符没有出现在f(i-1)长的子字符中，那么f(i) =  f(i-1) + 1;
        3. 如果第i个字符出现在f(i-1)的子字符中，那么
                3.1 计算第i个字符和它上次出现在字符串(整个完整的字符串，而不仅仅只是不出现重复字符的那个字串)中的位置的距离，记为d。
                3.2 如果d小于等于f(i-1)，那么f(i) = d;
                3.3 如果d大于f(i-1)，那么f(i) = f(i-1) + 1;
*/

#include <string>

int getLongestSubstringWithoutDup(std::string input) {
  int curLength = 0;
  int maxLength = 0;
  

  int lastIndex[26];  //  使用一个数组索引上一次字符出现的位置
  for (int i = 0; i < 26; i++) lastIndex[i] = -1;

  for (int i = 0; i < input.length(); i++) {
    int prevIndex = lastIndex[input[i] - 'a'];
    if (prevIndex < 0 || i - prevIndex > curLength) curLength++; // 2 和 3.3 
    else 
    {
      if (curLength > maxLength) maxLength = curLength;
      curLength = i - prevIndex;  // 3.2
    }
    lastIndex[input[i] - 'a'] = i;
  }

  if (curLength > maxLength) maxLength = curLength;  // 第26行补充
  return maxLength;
}