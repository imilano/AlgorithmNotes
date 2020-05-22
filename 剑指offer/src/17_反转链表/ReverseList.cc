// 题目描述
//         输入一个链表，反转链表后，输出新链表的表头。

// 题解:
//         递归法和非递归法




#include <iostream>

using namespace std;

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(nullptr) {
	}
};

class Solution {
public:
// 反转含义
//     ListNode* ReverseList(ListNode* pHead) {
//             if (pHead)
//                     return pHead;
//             ListNode *head = pHead;
//             ListNode *root = pHead;
//             int len = 0;
//             while(head) {
//                     head = head->next;
//                     len++;
//             };
//             if (len == 1)
//                     return pHead;
//             head = root;
//             int arr[len];
//             for (int i = 0,j=len-1; i < len,j>=0;i++,j--) {
//                     root->val = arr[j];
//                     root = root->next;
//             }
//             return head;
//     }
    ListNode* ReverseList(ListNode* pHead) {
            if (pHead == nullptr || pHead->next == nullptr)
                    return pHead;
            ListNode *current = pHead;
            ListNode *pre = nullptr, *next = nullptr;
            ListNode *head;
            while (current)
            {
                    next = current->next; // 关键一步，需要在前
                    if(next == nullptr) // 当current指向最后一个节点时,要把 head指向current
                            head = current;
                    current->next = pre; // 改变指针
                   // next->next = current;
                    pre = current;
                    current = next;
            }
            return head;
    }

    ListNode* ReverseList_recursive(ListNode* pHead) {
        //如果链表为空或者链表中只有一个元素
        if(pHead==nullptr||pHead->next==nullptr) return pHead;
         
        //先反转后面的链表，走到链表的末端结点
        ListNode* pReverseNode=ReverseList(pHead->next);// 最后pHead其实指向倒数第二个节点
         
        //再将当前节点设置为后面节点的后续节点
        pHead->next->next=pHead; // 倒数第二个节点和倒数第一个节点之间进行反转
        pHead->next=nullptr; // 让倒数第二个节点指向null
         
        return pReverseNode; // 然后递归的以两个节点为单位向前进行处理
         
    }
};