package array;


/**
 * @author Charlie
 * @When
 * @Description Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
 * <p>
 * The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
 * <p>
 * Note:
 * <p>
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may not use the same element twice.
 * @Detail 1. 最esay的方式,直接for循环暴力遍历即可
 * 2. 借鉴网上思路,采取2个指针的方式类似于快速排序,一个指针从左侧开始,另外一个指针从右侧开始
 * 相加太大,则右侧往小的移动,否则左侧往大的移动,因此关键是 该数组是有序的
 * @Attention: 1. 升序
 * 2. 返回的下标以1开始
 * 3. 找到第一次返回的值即可直接return
 *
 * 重点在于有序
 * @Date 创建时间：2020-02-18 16:11
 */
public class Array_167_TwoSumII_Inputarrayissorted
{
    public int[] twoSum(int[] numbers, int target)
    {
        int[] result = new int[2];
        for (int i = 0; i < numbers.length; i++)
        {
            for (int j = i + 1; j < numbers.length; j++)
            {
                if (numbers[i] + numbers[j] == target)
                {
                    result[0] = i + 1;
                    result[1] = j + 1;
                    return result;
                }

            }
        }

        return result;
    }

    public int[] twoSum2(int[] numbers, int target)
    {
        int[] result = new int[2];
        int left = 0;
        int right = numbers.length - 1;
        while (left < right)
        {
            int value = numbers[left] + numbers[right];
            if (value == target)
            {
                result[0] = left + 1;
                result[1] = right + 1;
                return result;
            } else if (value < target)
            {
                left++;
            } else
            {
                right--;
            }
        }


        return result;
    }


}
