/*    题目描述：数字以012345678910111213141516171819...的格式序列化到一个字符序列中。在这个序列中，第5位（从0开始计数）是5,第13位是1,第19位是4,等等。
 *              请写一个函数，求任意第n位对应的数字。
 * 
 * 求解，假设我们需要找第1001位置：
 *      根据规律可以发现，从0-9有10位，1001>10，所以不可能在一位数里；接下来在剩下的991位置里找，从10-99有180个数字，991>180，所以也不可能在两位数里面。
 *       接下来在三位数100-999里面找，发现有2700位，而2700 > 881，所以数字一定在三位数里面。因为881 = 270*3+1,所以数字为100+270=370的中间一位。
 */
#include <cmath>


int countOfIntegers(int digits) { // 计算每个位数里有多少数字，例如 一位数有10个，二位数有90个，三位数有900个
  if (digits == 1) return 10;
  int count = (int)std::pow(10, digits - 1);
  return count * 9;
}

int beginNumber(int digits) { // digits位的数字中最开始的那个数字，例如1位数为0,2位数为10,三位数为100
  if (digits == 1) return 0;
  return (int)std::pow(10, digits - 1);
}

int digitAtIndex(int index, int digits) { // 找出那一位数字
        int number = beginNumber(digits) + index/digits; // 为哪一位数
        int indexFromRight = digits - index % digits; //  从右边数的哪一位
        for (int i = 0; i < indexFromRight;++i) {
          number /= 10;
        }
        return number % 10;
}

int digitAtIndex(int index) {
  if (index < 0) return -1;

  int digits = 1; // 表示位数，一位数还是两位数 
  while(true) {
    int numbers = countOfIntegers(digits);
    if (index < numbers * digits) return digitAtIndex(index, digits); // 返回实际的数字

    index -= digits * numbers;
    digits++;
  }
  return -1;
}


