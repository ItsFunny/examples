package array;

/**
 * @author Charlie
 * @When
 * @Description Given an array,
 * rotate the array to the right by k steps, where k is non-negative.
 * @Detail 既向右移动k个位置
 * 1. 参考 https://blog.csdn.net/biezhihua/article/details/79535021
 * 2. 数学题,参考https://leetcode.com/problems/rotate-array/discuss/54250/Easy-to-read-Java-solution
 * 通过反转的形式,对于 i+k<=n-1 的元素,他的下标为i+k
 * 而对于i+k>n-1 的元素,需要
 * @Attention:
 * @Date 创建时间：2020-02-19 16:19
 */
public class Array_189_Rotate_Array_背
{
    public void rotate(int[] nums, int k) {
        k %= nums.length;
        reverse(nums, 0, nums.length - 1);
        reverse(nums, 0, k - 1);
        reverse(nums, k, nums.length - 1);
    }

    public void reverse(int[] nums, int start, int end) {
        while (start < end) {
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }
}
