package com.basic.sort;

import com.basic.utils.CommonUtil;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 15:26
 */
public class PopSort
{
    public static void main(String[] args)
    {
        CommonUtil.show();
        popSort(CommonUtil.ARR);
        CommonUtil.show();
    }

    public static void popSort(Integer[] arr)
    {
        //  冒泡排序就是暴力遍历比较
        //  如果后者小于则直接进行更换即可
        for (int i = 0; i < arr.length; i++)
        {
            for (int j = i + 1; j < arr.length; j++)
            {
                if (arr[j] < arr[i])
                {
                    int temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }

            }
        }
    }

}
