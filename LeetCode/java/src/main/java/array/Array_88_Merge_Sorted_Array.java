package array;

/**
 * @author Charlie
 * @When
 * @Description Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 * 2个有序数组,再合并为一个有序数组
 * @Detail 1. 拿nums2中的值一个一个与nums1中的值比较,nums2 j下标对应的值,找到在nums1中大于它的最小值
 * 将nums2中j下标的值,插入到nums1中i下标,后面元素后移,因此又会产生一个效率的问题,所以不可以从0开始匹配,
 * 从m和n开始匹配,当发生移动的时候只需要移动一个元素即可
 * 1 为错误的思路
 * 错误的地方在于,少一个变量k来代替全局下标的移动, k代表的是拼接后的下标,[0,n+m-1]
 * 并且要
 * @Attention:
 * @Date 创建时间：2020-02-14 12:58
 */
public class Array_88_Merge_Sorted_Array
{
    public void merge(int[] nums1, int m, int[] nums2, int n)
    {
        int i = m - 1;
        int j = n - 1;
        // 代替整个下标去移动
        int k = n + m - 1;
        // 找到最大值,因为是有序数组,所以后面的值是最大的
        for (; i >= 0 && j >= 0; )
        {
            if (nums1[i] < nums2[j])
            {
                nums1[k--] = nums2[j--];
            } else
            {
                nums1[k--] = nums1[i--];
            }
        }
        // 这里仅需要判断 较小的数组是否达到了底部
        // 原因在于 对于结果的操作是 基于nums1的基础上的,
        // 特殊情况1: nums1的值全部大于nums2,意味着最后j肯定大于0 ,此时只需要11赋值即可
        // 特殊情况2: nums1的值全部小于nums2,意味着nums2的所有值是自动添加到nums1后的,**并不需要对i判断,因为
        // nums1和nums2都是有序数组**
        // 所以这里只需要判断 j是否等于0
        while (j >= 0)
        {
            nums1[k--] = nums2[j--];
        }
    }

}
