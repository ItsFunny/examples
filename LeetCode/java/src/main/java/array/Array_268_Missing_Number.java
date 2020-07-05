package array;

/**
 * @author Charlie
 * @When
 * @Description Given an array containing n distinct numbers
 * taken from 0, 1, 2, ..., n, find the one that is missing from the array.
 * 找出丢失的数字,空间复杂度要求O(1),意味着不可申请额外的内存
 * @Detail 1. 最简单的方式是for循环遍历获取最大值之后,再for循环一次相减,返回的就是缺失的值
 * @Attention:  这个n指的是数组的长度,有n个,则最大值为n即可
 * @Date 创建时间：2020-02-20 16:18
 */
public class Array_268_Missing_Number
{
    public int missingNumber(int[] nums)
    {
        int max = nums.length;
        max = ((1 + max) * max) >> 1;
        for (int i = 0; i < nums.length; i++)
        {
            max -= nums[i];
        }
        return max;

    }

}
