#include <vector>
#include <iostream>

int main() {
        std::vector<int> v;
        v.reserve(10);
        // std::cout << v.max_size();
        // std::vector<int>::const_iterator cit = v.cbegin();
        // std::vector<int>::iterator it = v.begin();
        // *it = 5;
        // v.push_back(1);
        v[0] = 1;
        int a = v.at[0];
        std::cout << a << std::endl;
        std:: cout << v[0] << std::endl;



        std::cout << "before shrink, size: "<< v.size() <<", capacity: "<< v.capacity() << std::endl;
        v.shrink_to_fit();
        std::cout << "after shrink, size: " << v.size() << ", capacity: "<< v.capacity() << std::endl;
        


        // std::cout << v.capacity() << std::endl;

        // // for (int i = 0; i < 11 ; i++) {
        // //   v.push_back(i);
        // // }

        // std::cout << v.capacity() << std::endl;

        //   for (std::vector<int>::iterator it = v.begin(); it != v.end(); it++) {
        //     std::cout << *it << " ";
        //   }

          return 0;
}