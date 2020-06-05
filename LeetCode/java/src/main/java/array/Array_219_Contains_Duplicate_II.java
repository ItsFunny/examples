package array;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
 * @Detail 既判断相同的元素, 下标绝对值最大为k
 * 1. 遍历+map保存值
 * 不同在于
 * @Attention:
 * @Date 创建时间：2020-02-20 15:48
 */
public class Array_219_Contains_Duplicate_II
{
    public boolean containsNearbyDuplicate(int[] nums, int k)
    {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++)
        {
            if (map.containsKey(nums[i]))
            {
                Integer index = map.get(nums[i]);
                if (i - index <= k)
                {
                    return true;
                } else if (i - index > k)
                {
                    map.put(nums[i], i);
                }

            } else
            {
                map.put(nums[i], i);
            }

        }
        return false;

    }
}
