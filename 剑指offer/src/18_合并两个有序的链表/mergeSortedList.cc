// 题目描述
//         输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

// 题解:
//         非递归方式:采用归并排序思想
//         递归方式:




struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(nullptr) {
	}
};
class Solution {
public:
//         非递归方式
    ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
    {
        
        // if (pHead1 == nullptr || pHead2 == nullptr) {
        //         if (pHead1 == nullptr && pHead2 == nullptr )
//                 return nullptr;
        if (pHead1 == nullptr)
                return pHead2;
        else if(pHead2 == nullptr)
                return pHead1;
        // }

        ListNode *head = pHead1->val <= pHead2->val ? pHead1 : pHead2;
        ListNode* p = head;
        if (pHead1->val <= pHead2->val)
                pHead1 = pHead1->next;
        else
                pHead2 = pHead2->next;

        int v1 = 0, v2 = 0;
        while (pHead1 != nullptr && pHead2 != nullptr)
        {
                v1 = pHead1->val;
                v2 = pHead2->val;
                
                if (v1 <=v2 )
                {
                        p->next = pHead1;
                        p = p->next;
                        pHead1 = pHead1->next; 
                } else if (v1 > v2) {
                        p->next = pHead2;
                        p = p->next;
                        pHead2 = pHead2->next;
                }

        }
                if (pHead1 == nullptr) {
                        p->next = pHead2;
                }
                if (pHead2 == nullptr) {
                        p->next = pHead1;
        }
        return head;
    }
        // 递归方式
    ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
    {
        if (pHead1 == nullptr || pHead2 == nullptr) {
        if (pHead1 == nullptr && pHead2 == nullptr )
                return nullptr;
        if (pHead1 == nullptr)
                return pHead2;
        else
                return pHead1;
        }

        ListNode *head = nullptr;
        if (pHead1->val <= pHead2->val) {
                head = pHead1;
                head->next = Merge(pHead1->next, pHead2);

        }else
        {
                head = pHead2;
                head->next = Merge(pHead1, pHead2->next);
        }
        return head;
    }
};