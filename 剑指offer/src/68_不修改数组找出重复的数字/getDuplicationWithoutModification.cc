// 题目描述：在一个长度为n+1的数组里的所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的数组。

// 题解：把1～n的数字从中间开始分为两部分，前一部分为1～m，后一部分为m+1～n。扫描数组，如果1～m范围内的数字超过m，那么重复 数字肯定位于该区间；否则，m+1～n范围内一定包含重复的数字。


int countNumber(const int* numbers, int length, int start, int end ) {
    if (numbers == nullptr) return 0;
    int count=0;
    for (int  i=0; i<length; i++) 
        if (numbers[i] >= start && numbers[i] <= end) ++count;
    return count;
}

int getDuplication(const int* numbers, int length) {
    if (numbers == nullptr || length <=0) return -1;

    int start =1;
    int end = length-1;

    while(end >= start) {
        int middle = start+((end-start)>>1);
        int count = countNumber(numbers, length,start,middle);

        if (end ==start) {
            if (count >1) return start;
            else break;
        }

        if (count > middle-start+1) {
            end=middle;
        }else 
        start=middle+1;
    }
    
}
