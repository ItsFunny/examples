package stack;

import com.sun.org.apache.bcel.internal.generic.ISUB;

import java.security.Key;

/**
 * @author Charlie
 * @When
 * @Description You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2.
 * Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
 * <p>
 * The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.
 * 1. 既 sum1 中的元素为sum2中的子元素
 * 2. 除元素需要大于sum1中的外并且sum2中的下标也需要大于sum1中的下标
 * 3. 并且这个元素需要为 大于这个元素的最小值
 * <p>
 * 并不是存储下标,而是直接存储这个值,也并非是大于这个元素的最小值,大于它的第一个值即可
 * @Detail 1. 最简单的方式是直接2个for循环一个一个匹配
 * 题目没搞懂
 * 存储的不是下标,是元素,但是这道题难度不应该是easy 而应该是medium
 * @Attention:
 * @Date 创建时间：2020-03-17 16:27
 */
public class Stack_496_Next_Greater_Element_I
{
    public static int[] nextGreaterElement(int[] nums1, int[] nums2)
    {
        int[] result = new int[nums1.length];
        boolean exist = false;
        int minIndex = -1;
        for (int i = 0; i < nums1.length; i++)
        {
            for (int j = i + 1; j < nums2.length; j++)
            {
                if (nums2[j] > nums1[i])
                {
                    exist = true;
                    if (minIndex == -1 || nums2[j] < nums2[minIndex])
                    {
                        minIndex = j;
                    }
                }
            }
            if (exist)
            {
                exist = false;
            }
            result[i] = minIndex;
            minIndex = -1;
        }

        return result;
    }


    public static int[] nextGreaterElement2(int[] nums1, int[] nums2)
    {
        int[] result = new int[nums1.length];
        boolean exist = false;
        for (int i = 0; i < nums1.length; i++)
        {
            for (int j = i ; j < nums2.length; j++)
            {
                if (nums2[j] > nums1[i])
                {
                    exist = true;
                    result[i] = nums2[j];
                    break;
                }
            }
            if (exist)
            {
                exist = false;
            } else
            {
                result[i] = -1;
            }
        }

        return result;
    }


    public static void main(String[] args)
    {
        Stack_496_Next_Greater_Element_I.nextGreaterElement(new int[]{4, 1, 2}, new int[]{1, 3, 4, 2});
    }
}
