// 题目描述
// 将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。 数值为0或者字符串不是一个合法的数值则返回0

//  边界条件：
// 数据上下 溢出  用long存储 强转之前判断是否溢出 没溢出 强转 溢出 返回0 
// 空字符串
// 只有正负号
// 有无正负号
// 错误标志输出 

#include<string>
#include <limits>


class Solution {
public:
 int StrToInt(std::string str) {
   long long res = 0;
   int n = str.size(), symbol = 1;  // symbol标志正负号
   if (!n) return 0;
   int start;
   for (start = 0; start < n; start++)  // 跳过空格
   {
     if (str[start] == '-' || str[start] == '+' ||
         ('0' <= str[start] && str[start] <= '9'))
       break;
  }

  if (str[start] == '-')  // 标记正负
   symbol = -1;

  for (int i = (str[start] == '-' || str[start] == '+') ? start+1 : start; i < n; ++i){
   if (!('0' <= str[i] && str[i] <= '9')) // 检查字符串范围是否合法
    return 0;
   //res=res*10+str[i]-'0';
   res = (res << 1) + (res << 3) + (str[i] & 0xf);//这里利用位运算提高计算效率 。(res << 1) + (res << 3) = res * 2 + res * 8 = res * 10 。 字符'0'到'9'的ascii值的低4个二进制位刚好就是0到9所以str[i]&0xf等于str[i]-'0
  }
  res = res * symbol;  //用long存储 强转之前判断是否溢出 没溢出 强转 溢出 返回0 
  if (res > INT32_MAX || res < INT32_MIN) return 0;
  else
    return (int)res;
 }
};