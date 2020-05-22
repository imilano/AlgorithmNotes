// 题目描述
//         输入一个链表，输出该链表中倒数第k个结点。

// 题解：
//         双指针法，第一个指针先走k-1步，第二个指针再从头开始走，当第一个指针走到末尾时，第二个指针刚好指向倒数第k个节点
//         为保障鲁棒性，需要注意：
//                 1. pListHead 为空的情况
//                 2. 输入的链表长度小于k的情况
//                 3. 输入的参数为0。由于k是无符号数，那么在for循环中k-1得到的将不是-1,而是0xFFFFFFFF，会造成程序崩溃。



#include <iostream>

using namespace std;

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(nullptr) {
}
};

class Solution{
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {
            if (pListHead == nullptr || k < 0)
                    return nullptr;

            ListNode *head = pListHead;
            ListNode *tail = pListHead;
            int len = k-1;

            // 测试k是否小于链表长度，快指针先走k-1步，到达第k个节点
            for (int i = 0; i < len;i++) {
                    if (tail -> next)
                            tail = tail->next;
                else
                        return nullptr;
            }

            while(tail -> next != nullptr) {
                    head = head->next;
                    tail = tail->next;

            }
               return head;
    }
};