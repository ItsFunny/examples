package array;

import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
 * You may assume that the array is non-empty and the majority element always exist in the array.
 * @Detail 返回出现次数最多的元素
 * 1. 最简单的方法还是for循环遍历,然后通过一个map辅助排序
 * 2. 排序,既然题目要求是肯定存在,且出现次数大于n>>1次,则中间的元素一定就是
 * 3. 借鉴的网上的思路,通过抵消的形式,不管元素的大小,存在的互相抵消,意味着最后剩下的肯定是最多的那个元素
 * @Attention:
 * @Date 创建时间：2020-02-18 16:34
 */
public class Array_169_MajorityElement
{
    public static int majorityElement(int[] nums)
    {
        Map<Integer, Integer> m = new HashMap<>();
        for (int num : nums)
        {
            if (!m.containsKey(num))
            {
                m.put(num, 1);
            } else
            {
                Integer integer = m.get(num);
                integer++;
                m.put(num, integer);
            }

        }

        Integer max = 0;
        Integer maxCount = 0;
        for (Integer k : m.keySet())
        {
            if (m.get(k) > maxCount)
            {
                max = k;
                maxCount = m.get(k);
            }
        }
        return max;
    }


    public static int majorityElement2(int[] nums)
    {


        return 0;
    }

    public static int majorityElement3(int[] nums)
    {
        int majorValue = nums[0];
        int count = 0;
        for (int i = 0; i < nums.length; i++)
        {
            if (count == 0)
            {
                majorValue = nums[i];
                count++;
            } else if (nums[i] == majorValue)
            {
                count++;
            } else
            {
                count--;
            }

        }
        return majorValue;
    }

    private void qSort(int[] nums)
    {
        int left = 0;
        int right = nums.length - 1;

    }


}
