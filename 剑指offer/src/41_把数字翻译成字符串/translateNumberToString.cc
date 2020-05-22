// 题目要求：给定一个数字，按照如下规则翻译成字符串：0翻译成“a”，1翻译成“b”...25翻译成“z”。一个数字有多种翻译可能，例如12258一共有5种，分别是bccfi，bwfi，bczi，mcfi，mzi。实现一个函数，用来计算一个数字有多少种不同的翻译方法。

// 自上而下，从最大的问题开始，递归 ：
//                      12258
//                    /       \
//               b+2258       m+258
//               /   \         /   \
//           bc+258 bw+58  mc+58  mz+8
//           /  \      \        \     \
//       bcc+58 bcz+8   bwf+8   mcf+8  mzi
//         /        \       \     \
//    bccf+8        bczi    bwfi   mcfi
//      /
//  bccfi
// 有很多子问题被多次计算，比如258被翻译成几种这个子问题就被计算了两次。

// 自下而上，动态规划，从最小的问题开始 ：
// f(r)表示以r为开始（r最小取0）到最右端所组成的数字能够翻译成字符串的种数。对于长度为n的数字，f(n)=0,f(n-1)=1,求f(0)。
// 递推公式为 f(r-2) = f(r-1)+g(r-2,r-1)*f(r)；
// 其中，如果r-2，r-1能够翻译成位于10-25之间的字符，则g(r-2,r-1)=1，否则为0。
// 因此，对于12258：
// f(5) = 0
// f(4) = 1
// f(3) = f(4)+0 = 1
// f(2) = f(3)+f(4) = 2
// f(1) = f(2)+f(3) = 3 
// f(0) = f(1)+f(2) = 5

#include <string>
#include <vector>

int getTranslatin(int number) {
  if (number < 0) return 0;
  std::string numberInString = std::to_string(number);

  return getTranslationCount(numberInString);
}

int getTranslationCount(std::string s) {
  int len = s.length();
  if (len == 1) return 1;  // 至少需要两个数字
  std::vector<int> dp(len);

//   arr[len - 1] = 0, arr[len - 2] = 1;
  dp[len - 1] = 0, dp[len - 2] = 1;
  int count;
  for (int i = len - 1; i >= 0; i--) {
    count = 0;
    if (i < len - 1)
      count = dp[i + 1];
    else
      count = 1;
  if  (i  < len-1) {
          if (s.substr(i,i+1) >= "10" && s.substr(i,i+1) <= "25") {
            if (i < len - 2) count += dp[i + 2];
            else
              count += 1;
          }
    }
    dp[i] = count;
  }
  return dp[0];
}


 int translation ( int number )
{
     if ( number < 0 ) return 0 ;
    std::string s = std::to_string ( number ); //把数字转成字符串，方便引用各位数字
     int n = s . size ();
     if ( n == 1 ) return 1 ; //因为通项公式中最少须有两个初始值
    
    std::vector < int > dp ( n );//一般dp中序号与实际的序号对应起来，决定分配n空间还是n+1空间
    dp [ 0 ] = 1 ; //递推公式需要两个初值
    
    int temp = ( s [ 0 ] - '0' )* 10 + s [ 1 ] - '0' ; //构造组合数
     if (temp>=10  && temp<=25 )
        dp [ 1 ] = 2 ;
     else
        dp [ 1 ] = 1 ;
 
     int g = 0 ;
     for ( int i = 2 ; i <= n - 1 ; i ++)
     {
         int temp = ( s [ i - 1 ] - '0' )* 10 + s [ i ] - '0' ;
         if ( temp >= 10   && temp <= 25 ) 
            g = 1 ;
         else 
            g = 0 ;
        dp [ i ] = dp [ i - 1 ] + g * dp [ i - 2 ]; //假设dp[0]=1,则从dp[2]开始满足通项公式
     }
     return dp [ n - 1 ];
} 
