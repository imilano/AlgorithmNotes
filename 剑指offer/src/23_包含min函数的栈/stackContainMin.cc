/* 题目描述
 *       定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。
 *      注意：保证测试中不会当栈为空的时候，对栈调用pop()或者min()或者top()方法。
 * 
 * 
 * 题解：
 *     
 *
 *       利用一个辅助栈来存放最小值。每入栈一次，就与辅助栈顶比较大小，如果小就入栈，如果大就入栈当前的辅助栈顶；当出栈时，辅助栈也要出栈，
 *       这种做法可以保证辅助栈顶一定都当前栈的最小值 
 *       栈  3，4，2，5，1
 *       辅助栈 3，2，1
 */

#include <stack>
using namespace std;

class Solution {
public:
    void push(int value) {
            data_s.push(value);
            if(min_s.empty() || value <= min_s.top())  // 如果当前进栈元素比min_s要小，那么就把该元素也压入最小栈
                    min_s.push(value);
    }
    void pop() {
        if (data_s.top() == min_s.top())  // 如果数据栈出栈元素和最小栈出栈元素相等，那么顺便把最小栈的栈顶元素也出栈，以保持一致性
                min_s.pop();
        data_s.pop();
    }
    int top() {
            return data_s.top();
    }
    int min() {
            return min_s.top();
    }

private:
        stack<int> data_s;
        stack<int> min_s;
};