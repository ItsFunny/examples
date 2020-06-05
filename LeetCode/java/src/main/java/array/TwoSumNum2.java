package array;

import sun.util.resources.cldr.uk.CurrencyNames_uk;

import java.util.HashMap;
import java.util.Map;

/**
 * @author joker
 * @When
 * @Description 两数求和
 * @Detail 1. 最简单的方式,直接暴力循环即可, 时间复杂度 O(n^2) ,空间复杂度O(1)
 * 2. 通过map来辅助,遍历 然后判断 map中是否存储了剩下的值
 * @date 创建时间：2019-11-26 18:32
 */
public class TwoSumNum2
{

    public static int[] forceMatch(int[] nums, int target)
    {
        if (nums == null)
        {
            return null;
        }
        int[] res = new int[2];
        for (int i = 0; i < nums.length; i++)
        {
            for (int j = i + 1; j < nums.length; j++)
            {
                if (nums[i] + nums[j] == target)
                {
                    res[0] = i;
                    res[1] = j;
                    return res;
                }
            }

        }

        return res;
    }

    public static int[] findByMap(int[] array, int target)
    {
        if (array == null)
        {
            return null;
        }
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        map.put(array[0], 0);
        for (int i = 1; i < array.length; i++)
        {
            if (map.containsKey(target - array[i]))
            {

                return new int[]{map.get(target - array[i]), i};
            } else
            {
                map.put(array[i], i);
            }
        }

        return null;
    }
}
