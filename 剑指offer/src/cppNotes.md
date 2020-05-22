1. 野指针
   
   在分配空间给指针之后，一定要用nullptr来判断一下是否成功了。然后在删除这个指针的时候，也要用nullptr来赋给指针，杜绝成为野指针！
   ```c++
        int* p = new int;
        if (p == nullptr) { // 分配空间给 指针,一定要先判断是否分配成功
                //doSomething;
        }

        // 删除p之后,一定要记得把p置为nullptr
        delete p;
        p = nullptr;

   ```

2. 指针的引用
   
    即指针传递只是传了一个地址copy, 在函数内部改变形参所指向的地址，不能改变原实参指向的地址，仅可以通过修改形参地址的内容，来达到修改实参内容的目的（原C语言中的通过指针来互换值小函数例子），所以如果想通过被调函数来修改原实参的地址或给重新分配一个对象都是不能完成的，只能使用双指针或指针引用

```c++
//----------------------------------------
void test(int *t)
{
　　int a=1;
　　t=&a;  // 并不能改变实际参数p的地址.p和t本质上有着不同的地址
　　cout<<t<<" "<<*t<<endl;
}

int main(void)
{
    int *p=NULL;
    test(p);
    if(p==NULL)
    cout<<"指针p为NULL"<<endl;
    system("pause");
    return 0;
}
//--------------------------------------
void test(int *&p)
{
　　int a=1;
　　p=&a;  // 同时修改了指针值以及指针指向的值
　　cout<<p<<" "<<*p<<endl;
}

int main(void)
{
    int *p=NULL;
    test(p);
    if(p!=NULL)
    cout<<"指针p不为NULL"<<endl;
    system("pause");
    return 0;
}
```

3. alg find函数
   
   algorithm头文件使用的find函数，原型为：

```c++
template<class InputIterator, class T>
  InputIterator find (InputIterator first, InputIterator last, const T& val)
{
  while (first!=last) {
    if (*first==val) return first;
    ++first;
  }
  return last;
}
```

使用方法为：
```c++
 if (find(v.begin(),v.end(),"stringDemo") == v.end()) { // not find,the element is not in v
     doSomething();
 } else { // find element
     doAnotherThing(); 
 }

```
4. partition算法
   可以用partiton一次性找出无序数组中的第k个有序数字；此外，第k个有序数字之前的数字都是比k小的数字，也就是说，还可以用它找出第k个小的数字