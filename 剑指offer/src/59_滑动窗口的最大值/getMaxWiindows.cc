 /* 题目描述
 *      给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 
 *      针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
 * 
 * 思路：一个滑动窗口可以看成一个队列，只需要在队列中找出最大值即可。但为了得到滑动窗口的最大值，队列序可以从两端删除元素，因此使用双端队列。
 *     原则：
 *     对新来的元素k，将其与双端队列中的元素相比较
 *     1）前面比k小的，直接移出队列（因为不再可能成为后面滑动窗口的最大值了!）,
 *     2）前面比k大的X，比较两者下标，判断X是否已不在窗口之内，不在了，直接移出队列
 *     队列的第一个元素是滑动窗口中的最大值
 */


#include <vector>
#include <queue>

 using std::vector;
 using std::deque;

 class Solution {
  public:
   vector<int> maxInWindows(const vector<int>& num, unsigned int size) {
     vector<int> res;
     deque<int> s;
     for (unsigned int i = 0; i < num.size(); ++i) {
       while ( s.size() && num[s.back()] <= num[i])  //从后面依次弹出队列中比当前num值小的元素，同时也能保证队列首元素为当前窗口最大值下标
         s.pop_back();

       while ( s.size() && i - s.front() + 1 > size)  //当当前窗口移出队首元素所在的位置，即队首元素坐标对应的num不在窗口中，需要弹出
         s.pop_front();
         
       s.push_back(i);  //把每次滑动的num下标加入队列
       if (size && i + 1 >= size)  //当滑动窗口首地址i大于等于size时才开始写入窗口最大值
         res.push_back(num[s.front()]);
     }
     return res;
   }
};