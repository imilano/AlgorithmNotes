// 题目描述
//     给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。

// 题解：
//     1. 使用不断链的方法,主要需解决三个问题：
//         1.1. 链表中是否含环
//         1.2. 若含环，如何找到环的入口节点
//         1.3. 如何统计环中含有的节点数目
//     下面依次来解决这三个问题：
//         对于问题1.1,可以使用双指针法。慢指针一次走一步，慢指针一次走两步，如果快指针追上慢指针，则说明存在环；如果快指针走到了链表的末尾（next指向nullptr）都没有追上第一个指针，那么链表就不包含环。
//         对于问题1.2,如何找到环的入口节点，还是使用双指针法。假设环中的节点有n个，那么快指针向前走n步，然后快慢指针一起走，当慢指针走到链表末尾时，快指针走到环入口。
//         对于问题1.3,仍然使用双指针法。由1.1可知，若存在环，则双指针必然相遇，并且相遇的节点就在环中，那么我们可以从该节点开始遍历计数，当再次走到该节点时，所得到的count就是环中节点数。
//
//      2. 使用断链法
//           两个指针,一个在前面,另一个紧挨着这个指针,在后面.两个指针同时向前移动,诶移动一次,前面的指针的next指向null.
//           也就是说,访问过的节点都断开,最后到达的节点一定是尾节点的下一个,也就是环的第一个,此时已经是第二次访问环的地一个节点了,第一次访问的时候我们已经让他指向了null,故到此结束.


struct ListNode {
    int val;
    struct ListNode *next;
    ListNode(int x) :
        val(x), next(nullptr) {
    }
};

class Solution {
public:
    // 假设该链表带有一个头指针
    // 返回环的头节点
    ListNode* meetingNode(ListNode* head) {
        if ( head ==  nullptr) return head;
        ListNode* slow = head->next;
        if (slow == nullptr)return nullptr;

        ListNode* fast = slow->next;

        while (fast != nullptr && slow != nullptr)
        {
            if (fast == slow) return fast; // 如果存在环,则返回环
            slow = slow->next;
            fast = fast->next;

            if (!fast) fast = fast->next; // 注意判空
        }

        return nullptr;// 如果不存在,返回nullptr
        
    }

    // 不断链
    ListNode* EntryNodeOfLoop_v1(ListNode* pHead)
    {
        ListNode*  meetNode = meetingNode(pHead);
        if (meetNode == nullptr) return nullptr;

        // 得到环中节点数目
        int nodesInLoop = 1;
        ListNode* fast = meetNode;
        while(fast->next != meetNode) {
            fast = fast->next;
            nodesInLoop++;
        }

        // 找到环节点入口
        // 快节点向前走n步
        fast = pHead;
        for (int i=0; i< nodesInLoop;i++) {
            fast = fast->next;
        }

        // 慢节点再和快节点一起走
        ListNode* slow = pHead;
        while(slow != fast) {
            slow = slow->next;
            fast = fast->next;
        }

        return slow;
    }

        // 断链法
      ListNode* EntryNodeOfLoop_v2(ListNode* pHead) {
          if (pHead == nullptr) return nullptr;
          ListNode* pre = pHead;
          ListNode* cur = pHead->next;

          while(cur != nullptr) {
              pre->next = nullptr;
              pre = cur;
              cur = cur->next;
          }
          return pre;
      }
};