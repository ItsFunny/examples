package com.basic.sort;

import com.basic.utils.CommonUtil;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 15:58
 */

/*
    总结: 归并排序就是先分在治
    分:每次都一半一半来分,最终分成只有2个就可以开始归并了
    治:则是将2个数组从小到大插入到新的数组中,最后再讲新的数组的值复制到旧的数组,最终旧的数组就是有序的了
 */
public class MergeSortAlgorithm
{
    public static void main(String[] args)
    {
        CommonUtil.show();
        mergeSort(CommonUtil.ARR);
        CommonUtil.show();
    }

    public static void mergeSort(Integer[] arr)
    {
        // 归并算法分为2个步骤: 分治法  分 + 治

        Integer[] tempArr = new Integer[arr.length];

        mergeSort(arr, 0, arr.length - 1, tempArr);

    }
    // 分代表着将数组分为直至相邻的若干个小数组 ,直至分到了最小情况(既同一下标了),所以最小的数组的长度为2[5,6]
    // 为什么这样子就可以了呢,答案在于merge中,merge会遍历数组,判断大小,按顺序写入到临时队列中

    public static void mergeSort(Integer[] arr, int left, int right, Integer[] temp)
    {
        if (left < right)
        {
            int mid = (left + right) >> 1;
            // 对左边进行分
            mergeSort(arr, left, mid, temp);
            // 对右边进行分
            mergeSort(arr, mid + 1, right, temp);
            // 对左右进行治
            merge(arr, left, mid, right, temp);
        }
    }

    // 治的逻辑::
    // 对左边,右边进行遍历,同时会进行判断,选取小的值放入到临时数组中
    // 之后再进行遍历,此时只会遍历一边
    // 最后则是将临时数组中的元素复制到元数组中
    // 关键点在于: 要建立临时的变量,代替下标去移动
    public static void merge(Integer[] arr, int left, int mid, int right, Integer[] temp)
    {
        int i = left;
        int j = mid+1;
        int k = 0;
        // 进行归并

        while (i <= mid && j <= right)
        {
            if (arr[i] < arr[j])
            {
                temp[k++] = arr[i++];
            } else
            {
                temp[k++] = arr[j++];
            }

        }
        // 对数组中剩余的进行复制
        while (i <=mid)
        {
            temp[k++] = arr[i++];
        }
        while (j <=right)
        {
            temp[k++] = arr[j++];
        }
        k = 0;

        while (left <= right)
        {

            arr[left++] = temp[k++];
        }
    }

}
