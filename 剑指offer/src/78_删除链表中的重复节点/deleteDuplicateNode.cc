// 题目描述
//     在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。 例如，链表1->2->3->3->4->4->5 处理后为 1->2->5

// 题解：  1. 首先添加一个头节点，以方便碰到第一个，第二个节点就相同的情况
//         2.设置 pre ，last 指针， pre指针指向当前确定不重复的那个节点，而last指针相当于工作指针，一直往后面搜索

struct ListNode
{
        int val;
        ListNode *next;
        ListNode(int x) : next(nullptr) {}
 };


class Solution {
public:
    ListNode* deleteDuplication(ListNode* pHead)
    {
        if (pHead == nullptr || pHead->next == nullptr) return pHead;
        

        ListNode* head = new ListNode(0); // 只写成ListNode* head;不做初始化无法通过
        head->next = pHead;
        ListNode* pre = head;  // 注意这一步，指向新的头节点是为了能够较好处理只有两个相同节点的情况
        ListNode* current = pHead;

        while(current != nullptr) {
            if (current->next !=  nullptr && current->val == current->next->val) {
                while(current->next!= nullptr && current->val == current->next->val) current=current->next; // 一直向后搜索到重复节点的最后一个节点
                pre->next = current->next;
                // pre=current;
                current = current->next;
            }else{
                pre = pre->next;
                current = current->next;
            }
        }
        return head->next; 
    }
};