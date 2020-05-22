// 题目描述
// 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

// 题解：使用一个栈作为辅助栈，假设s1和s2,当压入元素时，把它压入S1中，当弹出元素时，如果s2不为空，直接弹出s2栈顶元素，如果s2为空，则将s1中的元素逐个压入s2,最后将s2的栈顶元素弹出。



#include <iostream>
#include <stack>

using namespace std;

class Solution
{
public:
    void push(int node) {
            stack1.push(node);
        }

    int pop() {
            
            if (stack2.empty()) {
                    while(!stack1.empty()) {
                            int tmp = stack1.top();
                            stack2.push(tmp);
                            stack1.pop();
                    }
            }

                int head = stack2.top();
                stack2.pop();
                return head;


        // solution 2
        //int topNum;
        //     if (!stack2.empty())
        //     {
        //             topNum = stack2.top();
        //             stack2.pop();
        //           //  return topNum;
        //     }else {
        //             while (!stack1.empty()) {
        //                     stack2.push(stack1.top());
        //                     stack1.pop();
        //             }
        //             topNum = stack2.top();
        //             stack2.pop();
        //     }
        //     return topNum;
    }

private:
    stack<int> stack1;
    stack<int> stack2;
};