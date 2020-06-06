package com.basic.sort;

import com.basic.utils.CommonUtil;

/**
 * @author joker
 * @When
 * @Description 简单选择排序
 * @Detail
 * @date 创建时间：2019-01-12 14:41
 */
public class SimpleSelectionSort
{
    public static void main(String[] args)
    {
        CommonUtil.show();
        simpleSelectionSort(CommonUtil.ARR);
        CommonUtil.show();
    }

    public static void simpleSelectionSort(Integer[] arr)
    {
        // 简单选择排序的核心就是0号放的是最小的元素,和1号放的是次小的元素,意味着需要暴力遍历
        for (int i = 0; i < arr.length; i++)
        {
            int min = arr[i];
            int pos = i;
            for (int j = i + 1; j < arr.length; j++)
            {
                if (arr[j] < min)
                {
                    min = arr[j];   //将最小的这个给min
                    pos = j;        // 记录最小的下标,方便更换
                }
            }
            arr[pos] = arr[i];
            arr[i] = min;

        }
    }
}
