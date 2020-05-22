```c++
        1. map简介。map<key,value>，每个key只能出现一次，map内部 自建一颗带有自动排序功能的红黑树(非严格意义上的二叉平衡树)，所以map内所有的数据都是有序的。由于使用了红黑树，所以map查找的复杂度是log(N)。
   
        2. 数据插入。map插入数据有三种方式
          - 用insert插入pair数据
                map<string,int> m;
                m.insert(pair<string, int>("name",12));
          - 使用insert插入value_type数据
                map<string, int> m;
                m.insert(map<string,int>::value_type("name",1));
          - 用数组方式插入数据
                map<sting,int> m;
                m["name"] = 1;
        // 前两种插入方式在效果上是一样的，用insert插入数据，当map中有这个关键字时，insert是插入不了数据的，但是用数组（第三种方式）的话，他尅覆盖以前该关键字对应的值：
        m.insert(map<string,int>::value_type("job",11));
        m.insert(map<string,int>::value_type("job",12)); // 并不会生效

        3. map遍历
         map遍历有三种方式：应用前向迭代器；应用反向迭代器（注意迭代器++）；用数组的形式（注意下标从1开始）




```