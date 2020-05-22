
```c++

## 1.定义vector

// vector 是一个变长数组，其定义如下：
vector<typename> name;

//typename 可以是任意类型，如果typename是容器，那么要在两个尖括号之间留一个空格，否则低版本的编译器容易把两个尖括号解析为移位运算：
std::vector<std::vector<int> > name; // 留空格


//vector 数组定义如下
std::vector<int> arr[n]; // 长度为n的vector数组

// vector的元素也可以是结构体，但是结构体要定义为全局变量

## 2. 访问vector

// 1. 通过下标访问，例如v[i]
// 2. 通过迭代器访问
std::vector<int>::iterator it = v.begin();
for (int i=0; i < len; i++) {
        std::cout << *(it+i) << std::endl;  // 实际上可以看出，迭代器就是指针的封装
}
// 迭代器实现了自增自减操作（++/--），所以另一种使用迭代器遍历的方式是：
for(std::vector<int>::iterator it = data.begin(); it != data.end(); it++) {  // 需要注意的是，美国人习惯左闭右开，所以begin指向第一个元素，end只想最后一个元素的后一个位置
        std::cout << *it<< std::endl;
}

// 注意，在常用STL中，只有在vector和string中，才允许使用v.begin()+3这样的写法。

## 3. vector常用函数
 // 1. v.push_back(datatype)，在后面添加一个元素，时间复杂度O(1)。
 // 2. 同理，v.pop_back(),从后面删除一个元素，时间复杂度O(1)
 // 3. v.size(),返回unsigned int，时间复杂度为O(1)。
 // 4. v.clear()， 清空v中的所有元素，时间复杂度O(N)，N为元素个数。
 // 5. v.insert。
 //     5.1 v.insert(it,data) 向迭代器it处插入一个元素data，时间复杂度O(N).**注意，不能用insert初始化元素**，下面这样的形式是错误的：
        v.insert(v.begin()+2, -1);   //正确方法
        
        for(int i=0;i<n;i++) {
                v.insert(i);  //  错误写法，无法实现初始化
        }

 //      5.2 v.insert(it, size,data); 在it位置插入size个data
 //      5.3 v.insert(it,it_first, it_second);  在it这个位置插入另一个容器的it_first到it_second之间的一段元素。
 // 6. v.erase()。erase有两种用法，一个是删除单个元素，第二个是删除一个区间之内的元素。
        v.erase(v.begin()); // 删除begin处的元素
        v.erase(v.begin(),v.begin()+2); // 删除[begin,begin+2)之间的元素，注意左闭右开
 // 7. v.size() 和 v.capacity()。vector 是自动增长的数组，vector的元素以连续方式存放，每一个元素都紧挨着前一个元素存储。vector进行内存分配时，其实际分配的容量要比当前所需空间多一些，其中size指的是当前容器真实占用的元素大小，capacity是指发生在realloc前能允许的最大元素个数，即预分配的内存空间。

 // 内存空间会自动增长，每当vector容器不得不分配新的存储空间时，会以加倍当前容量你的分配策略实现重新分配。
 
 // capacity为什么重要呢？容器的大小一旦超过capacity的大小，vector就会重新配置内部的存储器，导致和vector相关的reference、pointer和iterator都会失效。

 // vector的size和capacity属性对应两个方法，resize和reserve。使用resize(int)方法，容器内的对象空间是真正存在的。使用reserve(int)仅仅只是修改了capacity的值，容器内的对象并没有真实的内存空间。（此时使用[]运算符访问容器内的对象时，很可能相互先数组越界的错误。）

 // 由于vector的内存占用空间只增不减，比如你首先分配了10000个字节，然后erase掉后面9999个，留下一个有效元素，但是内存占用仍然为10000个，所有内存空间是在vector被析构之后才能被系统回收。如果需要空间动态缩小，可以考虑使用deque。
 
 // 避免内存重新分配的方法如下。方法一，使用reserve函数，在创建容器后，第一时间为容器分配足够大的空间，避免重新分配内存；方法二，利用构造函数构造出足够空间。

 // 另外，在STL中，拥有capacity属性的容器只有vector和 string。

 // 8.v.assign。对vector进行赋值，有两个构造函数
 //     8.1 v.assign(cpnst_iterator first, const_iterator last),将(first,last)区间的元素 赋值到当前的vector容器中
 //     8.2 v.assign(size_type n, const T&x).赋n个值为x的元素到vector中，这个容器会清除掉vectorr容器中以前的内容。



## 4. 算法

 // 1. reverse. reverse是algorithm头文件里的函数，使用方法为:
         reverse(v.begin(),v.end())
 // 2. sort. sort是头文件algorithm内的函数，可以使用sort对容器进行排序，默认是升高序排列，可以自定义cmp函数：
        sort(v.begin(),v.end());
        sort(v.begin(),v.end(),cmp);
        bool cmp(const int &a, const int &b) {
                return a >b; // 大的排在前面，降序排列
        }



```