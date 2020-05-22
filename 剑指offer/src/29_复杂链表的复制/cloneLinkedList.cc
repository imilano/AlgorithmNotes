/*  题目描述
 *    输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针指向任意一个节点），返回结果为复制后复杂链表的head。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
 *
 * 
 * 
 *  题解:O(n)的时间复杂度，O(1)的空间复杂度
 *      整个过程分为三步：
 *              1. 克隆主节点。复制原始链表的复杂节点N并创建新节点N'，然后把N'链接到N后面。
 *              2. 处理从节点。如果原始链表上的节点N的m_pSibling指向S，则它对应额复制节点N'的m_pSibling指向S的复制节点S'。
 *              3. 拆分链表。把第二步得到的链表拆分成两个链表，奇数位置上的节点组成原始链表，偶数位置上的及诶但组成复制出来的链表。
 */


struct RandomListNode {
    int label;
    struct RandomListNode *next, *random;
    RandomListNode(int x) :
            label(x), next(nullptr), random(nullptr) {
    }
};

// 三步法

class Solution {
public:
    // 复制主节点
 void CloneNodes(RandomListNode* head) { 
        RandomListNode* node = head;
        while (node != nullptr) {
          RandomListNode* clone = new RandomListNode(node->label);
          clone->next = node->next;
          clone->random = nullptr; // 此处如此设置是为了方便第二步处理

          node->next = clone;
          node = clone->next;
        }
 }
    // 复制兄弟节点
 void ConnectSiblingNodes(RandomListNode* head) {
        RandomListNode* node = head;
        
        while (node != nullptr) {
               RandomListNode* clone = node->next;
                if (node->random != nullptr) {
                  clone->random = node->random->next;
                }
                node = clone->next;
                // clone = node->next;
        }
 }

// 从一个链表分离出两个链表需要三个指针：
//         nhead 用于指向原节点 
//         chead 用于指向克隆节点 
//         result 用于表示连接出来的克隆节点

//         nhead和chead都需要不断的向前遍历

// 分离两个链表
 RandomListNode* ReconnectNodes(RandomListNode* head) {
   RandomListNode* nhead = head;
   RandomListNode* chead = nullptr;
   RandomListNode* result = chead;

   if (head != nullptr)
   {
     result = chead = nhead->next;
     nhead->next = chead->next;
     nhead = chead->next;
   }
   while (nhead != nullptr) {
     chead->next = nhead->next;  // chead向前遍历
     chead = chead->next;
     nhead->next = chead->next;  // nhead向前遍历
     nhead = nhead->next;
   }
   return result;
 }

 RandomListNode* Clone(RandomListNode* pHead) {
   CloneNodes(pHead);
   ConnectSiblingNodes(pHead);
   return ReconnectNodes(pHead);
 }
};

// 递归法
class SolutionRec {
public:
    RandomListNode* Clone(RandomListNode* pHead)
    {
        if(pHead==nullptr)
            return nullptr;
         
        //开辟一个新节点
        RandomListNode* pClonedHead = new RandomListNode(pHead->label);
        pClonedHead->next = pHead->next; // 处理最后一个节点
        pClonedHead->random = pHead->random;
         
        //递归其他节点
        pClonedHead->next=Clone(pHead->next);
         
        return pClonedHead;
    }
};