/* 题目描述
 *       输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）
 *
 * 题解：
 *      建立一个辅助栈，把输入的第一个序列中的数字依次压入该辅助栈，并按照第二个序列的顺序依次从该栈中弹出数字。具体过程如下：
 *              - 如果下一个弹出的数字刚好是栈顶数字，那么直接弹出；
 *              - 如果下一个弹出的数字不在栈顶，则把压栈序列中还没有入栈的数字压入辅助栈，直到把下一个需要弹出的数字压入栈顶为止；
 *              - 如果所有数字都压入栈后仍然没有找到下一个弹出的数字，那么该序列不可能是一个弹出序列。
 *      
 */
#include <vector>
#include<stack>

using namespace std;

class Solution {
public:
    bool IsPopOrder(vector<int> pushV,vector<int> popV) {
        if (pushV.empty() && popV.empty())
                return false;
        stack<int> s; // 辅助栈
        int index = 0; // 标记出栈序列的位置
        int len = pushV.size();

        for (int i = 0; i < len;i++) {
                s.push(pushV[i]);
                while(!s.empty() && s.top() == popV[index]) { 
                        //如果栈顶元素和出栈序列元素相等，则将该栈顶元素出栈，并且增加出栈序列的下标
                        s.pop();
                        index++;
                }
        }

        return s.empty(); // 所有入栈元素都已入栈完成，和出栈元素匹配的元素都已经出栈，若栈中还有元素，则不匹配。
    }
}; 