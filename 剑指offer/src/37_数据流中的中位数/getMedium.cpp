// 题目描述：
//        如何得到一个数据流中的中位数？如果从数据流中读出技术个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间爱你两个数的平均值。

// 题解：
//     1. 使用AVL树。AVL树是一棵平衡二叉树，我们更改AVL树的平衡因子是左右子树的节点数之差。有了这个改动，我们就可以在O(nlogn)的时间内插入一个节点，并且在O(1)时间内得到所有节点的中位数。
//     2. 使用最大最小堆。通过观察可以发现，容器中数据被中间的一个或两个数据分隔成两部分。位于左边部分的数据都比右边部分的数据小，使用两个指针p1、p2,p1指向的数据是左边部分的最大数，p2指向的数据
//         是右边部分的最小数。可以使用最大堆实现左边的数，最小堆实现右边的数。
//         往堆中插入一个数据时，要保证数据能够平均分配到两个堆中，因此两个堆中的数据数目之和不能超过1.为了实现平均分配，可以在数据的总数目为偶数时把数据插入最小堆，否则插入最大堆。
//         还要保证最大堆中的所有数据都要小于最小堆中的数据。当数据的总数是偶数时，按照规则需要将其插入最小堆，但是如果该数字比左边的最大堆的最大值要小呢？此时应该将其插入左边的最大堆，然后将左边最大堆的最大值插入最小堆。
//         当需要把一个数据插入最大堆，但是这个数据小于最小堆中的一些数据时，应该把该数据插入最小堆，然后把最小堆的堆顶元素插入最大堆。


#include <vector>
#include <algorithm>

class Solution {
public:
    void Insert(int num)
    {
        if((min.size()+max.size())&1 == 0) {
                // 为保证平均分配到两个堆，偶数个分配到最小堆（右侧），奇数个分配到最大堆（左侧）
                if (max.size()>0 && num < max[0]) { // 如果比最大堆（左侧）的最大元素都要小，则将该元素插入最大堆，然后将最大堆的最大元素插入最小堆（右侧）
                  max.push_back(num);
                  std::push_heap(max.begin(), max.end(), std::less<int>());

                  num = max[0];
                  // pop_heap()是在堆的基础上，弹出堆顶元素。这里需要注意的是，pop_heap()并没有删除元素，而是将堆顶元素和数组最后一个元素进行了替换，如果要删除这个元素，还需要对数组进行pop_back()操作。
                  // push_heap()：新添加一个元素在末尾，然后重新调整堆序。也就是把元素添加在底层vector的end()处。
                  std::pop_heap(max.begin(), max.end(), std::less<int>());
                  max.pop_back();
                }
                min.push_back(num);
                std::push_heap(min.begin(), min.end(), std::greater<int>());
        }else {
                if (min.size() > 0 && num >  min[0]) { // 如果比最小堆（右侧）的元素要小，那么把他插入左边元素，然后将左边元素的最大值插入右边
                  min.push_back(num);
                  std::push_heap(min.begin(), min.end(), std::greater<int>());

                  num = min[0];

                  std::pop_heap(min.begin(), min.end(), std::greater<int>());
                  min.pop_back();
                }
                max.push_back(num);
                std::push_heap(max.begin(), max.end(), std::less<int>());
        }
    }

    double GetMedian() { 
        int size = min.size() + max.size();
        double median = 0;
        if ((size & 1) == 0) median = (min[0] + max[0]) / 2; // 如果总数是偶数，那么取中间元素平均值
        else
          median = min[0]; //  如果总数是奇数，那么取最小堆的最小值（跟入堆顺序相关）
        return median;
    }

   private:
    std::vector<int> min;
    std::vector<int> max;
};