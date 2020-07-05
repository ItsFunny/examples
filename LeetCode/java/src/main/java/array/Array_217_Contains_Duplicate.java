package array;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Given an array of integers, find if the array contains any duplicates.
 * <p>
 * Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
 * 如果有重复的元素则返回true
 * @Detail 1. 最简单的方式是for 循环加上一个map即可
 * 2. 通过sort即可,判断前后是否相等即可
 * @Attention:
 * @Date 创建时间：2020-02-19 17:12
 */
public class Array_217_Contains_Duplicate
{
    public boolean containsDuplicate(int[] nums)
    {
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums)
        {
            if (map.containsKey(num))
            {
                return true;
            } else
            {
                map.put(num, 1);
            }

        }
        return false;
    }


}
