/*
 * c++ 中有两个地方定义getline函数.一个是在各种stream中,例如std::iostream::getline、std::fstream::getline、std::istream::getline等。
 * 例如 cin.getline(char* s, streamsize n)或者cin.getline(char* s,streamsize n,char delim)，注意，这里的getline读入空白符，但是不包括最后的换行符
 * 
 * 另一个实在std空间的全局 函数，因为这个getline函数的参数使用了string字符串，所以声明在<string>头文件中。此处的getline也是不读入换行符的
 */


#include <iostream>



int main() {
  char* str;
  int len = 10;
//   std::cin.getline(str, len);

  std::string result;
  std::getline(std::cin, result);
  std::cout << result;
  return 0;
}