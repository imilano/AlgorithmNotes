// 输入两个链表，找出它们的第一个公共结点。
#include <algorithm>

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(nullptr) {
	}
};

// 公共节点只会出现在尾部，因而如果有公共节点，那么两个链表一定是Y状。如果从两个链表的尾部开始往前进行比较，那么最后一个相同的节点就是我们要找的节点。可是单向链表中只能从前向后遍历，最后才能到达尾节点。最后到达的节点却要最先进行比较，是一个栈结构。
// 方法一：使用两个辅助栈，将节点入栈；这样两个链表的尾节点就位于两个栈的栈顶，接下来比较两个栈顶的节点是否相同即可。若相同，则弹出继续比较下一个，直到找到最后一个尾节点。适用于两链表长度相同的情况。
// 方法二：遍历两个链表，得到其长度m和n，假设m大于n，则第一个链表向前走m-n步。接下来两指针一起出发，直到找到他们第一个相同的节点。

class Solution {
public:
    ListNode* FindFirstCommonNode( ListNode* pHead1, ListNode* pHead2) {
      int len1 = 0, len2 = 0;
      if (pHead1 == nullptr || pHead2 == nullptr) return nullptr;

      ListNode* p1 = pHead1;
      ListNode* p2 = pHead2;

      while(p1 != nullptr) {
        len1++;
        p1 = p1->next;
      }
      while(p2 != nullptr) {
        len2++;
        p2 = p2->next;
      }

      int lessLen = std::min(len1, len2);
      p1 = pHead1;
      p2 = pHead2;

      if (lessLen == len1) {
        int d = len2 - len1;
        while(d>0) {
          p2 = p2->next;
          d--;
        }
      }

      if (lessLen == len2) {
        int d = len1 - len2;
        while(d>0) {
          p1 = p1->next;
          d--;
        }
      }

      while(p1 != nullptr && p2 != nullptr && p1 != p2) {
        p1 = p1->next;
        p2 = p2->next;
      }

      ListNode* result = p1;
      return result;
    }
};