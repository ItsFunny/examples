package com.basic.algorithm.timewheel.jianzhioffer;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * @author joker
 * @When
 * @Description 题目描述
 * 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
 * 使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
 * 并保证奇数和奇数，偶数和偶数之间的相对位置不变。
 * @Detail 直接采用list来做, 时间复杂度O(n)
 * 推荐参考链接https://www.nowcoder.com/questionTerminal/beb5aa231adc45b2a5dcc5b62c93f593
 * @date 创建时间：2019-05-25 22:21
 */
public class _13_re_order_array
{
    public void reOrderArray(int[] array)
    {
        List<Integer> list = new ArrayList<>();
        for (int i : array)
        {
            if (i % 2 == 1)
            {
                list.add(i);
            }
        }
        for (Integer i : array)
        {
            if (i % 2 == 0)
            {
                list.add(i);
            }

        }
        for (int i = 0; i < list.size(); i++)
        {
            array[i] = list.get(i);
        }
    }
}
