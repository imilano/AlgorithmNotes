#include <vector>
#include <cstdio>
#include <stack>

using namespace std;

struct ListNode
{
        int val;
        struct ListNode *next;
        ListNode(int x) : val(x), next(nullptr)
        {
        }
 };


/**
*  struct ListNode {
*        int val;
*        struct ListNode *next;
*        ListNode(int x) :
*              val(x), next(NULL) {
*        }
*  };
*/
class Solution {
public:

    // 使用栈
    std::vector<int> printListFromTialToHead_v1(ListNode* head) {
        vector<int> value;
        ListNode *p=nullptr;
        p=head;
        stack<int> stk;
        while(p!=nullptr){
            stk.push(p->val);
            p=p->next;
        }
        while(!stk.empty()){
            value.push_back(stk.top());
            stk.pop();
        }
        return value;
    }

    // 使用递归
    std::vector<int> result;
    std::vector<int> printListFromTialToHead_v2(ListNode* head) {
        if (head == nullptr) return;
        printListFromTialToHead_v2(head->next);
        result.push_back(head->val);
    }
    
    vector<int> printListFromTailToHead(ListNode* head) {

    }
};