// 题目描述：使用两个队列实现一个栈
// 题解：
//     入栈操作：
//             1. 当A、B队列均为空时，将元素进入A队列。
//             2. 当A、B队列有一个非空时，将元素压入非空的那个队列
//     出栈操作：
//             1. 当A、B队列均为空时，出栈为NULL
//             2. 当A、B队列有一个非空时，将不为空的队列（假设为B）中的前N-1个元素一依次入队到另一个队列（队列A）中，再将B队列中剩下的最后一个元素出队并返回。

#include <queue>
class Solution{
    private:
        std::queue<int> A;
        std::queue<int> B;
    public:
        void push(int n) {
            // 如果B不空，则压入B
            if (!B.empty()) {
                B.push(n);
                return;
            }
            // 如果A和B都空，或者A不空，则压入A
            A.push(n);
            return;
        }

        int pop() {
            // 如果B不空，则将B的N-1个元素弹出压入A中，然后返回B中最后一个元素
            if (!B.empty()) {
                while(B.size() > 1) {
                    A.push(B.front());
                    B.pop();
                }
                int temp = B.front();
                B.pop();
                return temp;
            }
            if (A.empty()) { // 如果A也为空，返回错误
                return -1;
            }

            //如果B为空，A不为空，同理，将A的N-1个元素（如果有的话）压入B中，弹出最后一个元素
            while(A.size() > 1) {
                int temp = A.front();
                B.push(temp);
                A.pop();
            }
            int temp = A.front();
            A.pop();
            return temp;

        }

        bool isEmpty()  {
            return A.empty() && B.empty();
        }
};