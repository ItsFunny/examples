package array;

import java.util.*;

/**
 * @author joker
 * @When
 * @Description 去除重复的元素
 * @Detail 1. 最简单的方式,通过hashSet来去重
 * 2. 通过特性,这个数组是sorted的有序数组,意味着下一个index的值必定比当前值要大
 * @date 创建时间：2019-11-27 17:16
 * 收获: 充分从已知条件中获取有用信息,如有序数组,有序结合题意去除重复,既标准数据是 一组有序的不重复的,既下一位的值必定大于
 * 上一位的值,因此
 */
public class RemoveDuplicatesNum26
{

    public int removeDuplicates(Integer[] nums)
    {
        Integer i = 0;
        for (Integer num : nums)
        {
            // i==0 代表是第一个
            // 因为是有序数组,所以后面的肯定比前面的大
            if (i == 0 || num > nums[i - 1])
            {
                // 这步不怕覆盖老的数据吗
                // 并不会,因为i只能是>=原先的index
                nums[i++] = num;
            }

        }
        return i;
    }

}
