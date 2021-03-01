
// space complexity： O(1)， 因为是尾递归
// time complexity: O(logk)，因为每次k都减少一半，而k=(m+n)/2,所以每次
class Solution {
public double findMedianSortedArrays(int[] nums1, int[] nums2) {
    int length1 = nums1.length;
    int length2 = nums2.length;

    // 如果length1+length2是奇数，也就是一共有奇数个数，那么median左右两边的个数相等，那么left=right；
    // 如果是偶数，那么median是中间两个数的加权和，left 为左边的数的index，right为右边数的index
    int left= (length1+length2 +1)/2;
    int right = (length1 + length2 + 2)/2;

    return (getKth(nums1,0,length1-1,nums2,0,length2-1,left) + getKth(nums1,0,length1-1,0,length2-1,right)) /2;
}

    private int getKth(int[] nums1, int start1, int end1, int[] nums2, int start2, int end2, int k) {
        // 两个子数组的元素个数
        int len1 = end1 - start1  + 1;
        int len2 = end2 - start2 + 1;

        // 如果有子数组为空，那么一定是nums1为空
        if (len1 > len2) return getKth(nums2,start2,end2,nums1,start1,end1,k);

        // If left half is empty, then we just need to return kth number of number2
        if (len1 == 0) return nums2[start2 + k -1];

        // If we need to find the 1th number, we just need to return the smaller one of two arrays
        if (k == 1) return math.Min(nums1[start1],nums1[start2]);

        // In case that the number of array is less than k/2, if so, we just drop all the element on the left
        int i = start1 + math.Min(len1,k/2) -1;
        int j = start2 + math.Min(len2, k/2) -1;

        // divide, drop half of k element
        if (nums1[i] > nums2[j]) {
            return getKth(nums1,start1,end1,nums2,j+1,end2, k-(j - start2) +1);
        } else {
        return getKth(nums1,i+1, end1, nums2, start2,end2,k - (i - start1) + 1);
        }
    }
}