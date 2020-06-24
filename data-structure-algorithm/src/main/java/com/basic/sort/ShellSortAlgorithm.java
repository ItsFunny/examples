package com.basic.sort;

import com.basic.utils.CommonUtil;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 10:35
 */
public class ShellSortAlgorithm
{
    public static void main(String[] args)
    {
        CommonUtil.show(CommonUtil.ARR);
        ShellSortAlgorithm algorithm = new ShellSortAlgorithm();
        algorithm.shellSort(CommonUtil.ARR);
        CommonUtil.show(CommonUtil.ARR);
    }

    public void shellSort(Integer[] arr)
    {
        // 希尔排序是直接插入算法的优化:
        // 将一个数组分成多块,对每块进行插入排序
        // 直接插入排序其实就是步长为1的希尔排序
        int stride = arr.length;
        while (stride != 1)
        {
            stride >>= 1;
            // 对每个分组进行排序
            for (int i = 0; i < stride; i += stride)
            {
                for (int j = 0; j < arr.length; j += stride)
                {
                    int temp = arr[j];
                    int k = j - stride;
                    for (; k >= 0 && arr[k] >= temp; k -= stride)
                    {
                        arr[k + stride] = arr[k];
                    }
                    arr[k + stride] = temp;
                }
            }
        }
    }


}
