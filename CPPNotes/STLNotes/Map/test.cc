
    //第三种：用数组方式，程序说明如下  
      
    #include <map>  
      
    #include <string>  
      
    #include <iostream>  
      
    using namespace std;  
      
    int main()  
      
    {  
      
        map<int, string> mapStudent;  
      
        mapStudent.insert(pair<int, string>(1, "student_one"));  
      
        mapStudent.insert(pair<int, string>(2, "student_two"));  
      
        mapStudent.insert(pair<int, string>(3, "student_three"));  
      
        int nSize = mapStudent.size();  
      
    //此处应注意，应该是 for(int nindex = 1; nindex <= nSize; nindex++)  
    //而不是 for(int nindex = 0; nindex < nSize; nindex++)  
      
        for(int nindex = 1; nindex <= nSize; nindex++)  
      
            cout<<mapStudent[nindex]<<endl;  
      
    }  
