package array;

import java.util.Arrays;
import java.util.Collections;

/**
 * @author Charlie
 * @When
 * @DescriptionGiven a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
 * 在 int[] 中,找到第3最大的值,无则返回最大值,时间复杂度O(n)
 * @Detail 1. 排序直接返回index=2的元素
 * 2. 傻逼题,只有1+1=2的做法,直接定义3个变量即可,无任何参考价值
 * @Attention:
 * @Date 创建时间：2020-02-21 16:41
 */
public class Array_414_third_maximum_number
{
    public int thirdMax(int[] nums)
    {
        Integer[] largest = new Integer[3];
        for (int n : nums)
        {
            if (largest[0] == null || n > largest[0])
            {
                largest[2] = largest[1];
                largest[1] = largest[0];
                largest[0] = n;
            } else if (largest[0] != null && largest[0] > n && (largest[1] == null || n > largest[1]))
            {
                largest[2] = largest[1];
                largest[1] = n;
            } else if (largest[1] != null && largest[1] > n && (largest[2] == null || n > largest[2]))
            {
                largest[2] = n;
            }
        }
        return largest[2] == null ? largest[0] : largest[2];
    }
}
