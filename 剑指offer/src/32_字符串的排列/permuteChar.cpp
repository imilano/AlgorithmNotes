/*  题目描述：
 *     输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
 * 
 *  题解：递归思维。将一个字符串分为第一个字符和后面的字符，第一个位置的所有可能就是第一个位置和后面的所有位置进行交换。
 *      然后在对剩下的字符求其排列，求其排列的过程和上述是一样的，因而是一个递归的过程。详情可见后面的全排列图
 */


#include<vector>
#include <string>
#include<algorithm>
#include <iostream>

using namespace std;

class Solution {
public:
    vector<string> result;
    vector<string> Permutation(string str) {
      if (str.empty()) return result;
      permute(str, 0);
      sort(result.begin(), result.end());
      return result;
    }
    void swap(char& s, char& d) { 
        char temp = s;
        s = d;
        d = temp;
    }

    void permute(string s, int index) { // index 表示的是当前正在做哪一个字符的全排列
      if (index == s.size()-1) { // 必须加大括号，否则下面的if会和后面的else匹配，导致else匹配错误
        if (find(result.begin(),result.end(),s) == result.end()) // 去重作用
            result.push_back(s);
      }
      else {
        for (int i = index; i < s.size(); i++) {
          swap(s[index],s[i]); // 交换位置，进行全排
          permute(s, index + 1);
          swap(s[i], s[index]); // 交换结束之后还要交换回来
        }
        }
    }

    // 使用库函数，还会自动去重
void getPermutation(string s) {
  do {
    cout << s << endl;
  } while (next_permutation(s.begin(), s.end()));
}
};
