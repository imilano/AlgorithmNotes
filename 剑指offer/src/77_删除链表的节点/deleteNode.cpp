// 题目描述：在O(1)时间内删除链表节点。
//         给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间内删除该节点。

// 题解：删除一个节点不必从头开始查找。我们尅很方便地得到要删除的节点的下一节点。如果我们把下一节点的内容复制到需要删除的节点上覆盖原有的内容，再把下一节点删除，那么i就能实现在O(1)时间内删除节点。
//     在删除节点的过程中需要注意的三个情况：
//         - 如果链表只有一个节点，并且删除该节点————删除该节点，并且将头节点设置为nullptr
//         - 如果链表有多个节点，删除的节点位于链表中间 ———— 直接删除
//         - 如果链表有多个节点，删除的节点位于链表末尾 ———— 遍历到末尾，删除尾节点

struct ListNode
{
        int val;
        ListNode *next;
        ListNode(int x) : next(nullptr) {}
 };


class Solution {
    public:
        void deleteNode(ListNode** pListHead, ListNode* pToBeDeleted) {
            if (pToBeDeleted == nullptr || pListHead ==  nullptr) return;

            if(pToBeDeleted->next != nullptr) { // 如果要删除的节点不是尾节点
                ListNode* pNext = pToBeDeleted->next;
                pToBeDeleted->val = pNext->val;
                pToBeDeleted->next = pNext->next;

                delete pNext;
                pNext = nullptr;
            }else if (*pListHead == pToBeDeleted) { // 如果链表只有一个节点
                delete pToBeDeleted;  // 我们在删除一个指针之后，编译器只会释放该指针所指向的内存空间，而不会删除这个指针本身,因而在删除一个指针之后，还需要将其设置为nullptr
                pToBeDeleted = nullptr;
                *pListHead = nullptr;
            } else { // 如果删除的节点是尾节点：由于是尾节点，所以需要一次遍历到尾节点之前的节点
                ListNode* pNext = *pListHead;
                while(pNext->next != pToBeDeleted) pNext = pNext->next;  // 一直遍历到尾节点之前的节点
                pNext->next = nullptr;
                delete pToBeDeleted;
                pToBeDeleted = nullptr;
            }

        }

};