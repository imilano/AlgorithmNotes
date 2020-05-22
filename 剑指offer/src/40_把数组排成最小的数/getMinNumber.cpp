// 题目描述：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
// 将int转换为字符串，再将字符串拼接，按字典序比较


#include<vector>
#include <string>
#include <algorithm>
// #include <cstdlib>



class Solution {
public:
static bool cmp(std::string s1, std::string s2) {  // compare 函数必须为static
        std::string cat1="", cat2="";
        cat1 = s1 + s2;
        cat2 = s2 + s1;

        return cat1 < cat2;
}
 std::string PrintMinNumber(std::vector<int> numbers) {
        std::string result = "";
        std::vector<std::string> iTos;

        if (numbers.empty()) return result;

        for (std::vector<int>::iterator it = numbers.begin();
             it != numbers.end();it++) {
          iTos.push_back(std::to_string(*it));
        }

        std::sort(iTos.begin(), iTos.end(), cmp);

        for (std::vector<std::string>::iterator it = iTos.begin();
             it != iTos.end();it++) {
          result += *it;
        }

        return result;
 }
};