#include <map>
#include <string>
#include <iostream>
#include <algorithm>

int main() {
        std::map<std::string,std::string> name;

        // 检查是否插入成功，若成功，则insert_pair的第二个元素为true
        std::pair<std::map<std::string, std::string>::iterator, bool> insert_pair;

        // map插入方式一
        name.insert(std::pair<std::string, std::string>("job", "name"));
        insert_pair = name.insert(std::pair<std::string, std::string>("job", "nam"));

        // 检测是否插入成功
        if (insert_pair.second == false)
          std::cout << "insert failed" << std::endl;
        // map插入方式二
        name.insert(
            std::map<std::string, std::string>::value_type("name", "job"));

        // map遍历的三种方式
        name.clear();
        name["one"] = "one";
        name["two"] = "two";
        name["three"] = "three";


  std::map<std::string, std::string>::iterator it = name.begin();
  // 正向迭代器
  for (; it != name.end(); it++) {
    std::cout << it->first << ":" << it->second << std::endl;
  }

  // 反向迭代器
  std::map<std::string, std::string>::reverse_iterator rit = name.rbegin();

  for (; rit != name.rend(); rit++) {
    std::cout << rit->first << ":" << rit->second << std::endl;
  }
 
  // 数组方式 这样的方式要注意循环变量和原先的map里的key一致
//   for (int i = 0; i < name.size();i++) {
//     std::cout << name[i] << std::endl;
//   }
  return 0;
}