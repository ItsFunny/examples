package com.basic.sort;

import java.util.ArrayList;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-18 21:37
 */
/*
  总结:
    1. 先写paration函数,获取一个标准值的下标(通过先定义一个标准值,
            然后对左边和右边进行遍历匹配,满足则移动,不满足则更换位置)
    2.然后对左边进行排序,对右边进行排序
  注意点:
    1. 记得先写退出条件,在快速排序中因为涉及到递归对左右两边排序,既当left=right的时候就可以退出了
 */
public class QSort
{
    public void qSort(Integer[] arr)
    {
        qSort(arr, 0, arr.length - 1);
    }

    // qSort 的关键在于有一个标准值,左边的树都小于这个值右边的都大于这个值
    public void qSort(Integer[] arr, Integer start, Integer end)
    {
        if (start < end)
        {
            Integer paration = paration(arr, start, end);
            qSort(arr, start, paration);
            qSort(arr, paration + 1, end);
        }
    }

    // 既左边的小于这个值,右边的都大于这个值
    public Integer paration(Integer[] arr, Integer start, Integer end)
    {
        // 取一个标准值作为参考,然后递归进行比较,如果取左边的话,则从右边开始递归
        // 相反如果先取的右边,则先判断左边
        int stanard = arr[start];
        while (start < end)
        {
            // 右边的值都是大于左边的,因此一旦有值小于标准则,则需要将其换到左边去,同时这个时候左边的值是刚好
            // 是临界值,既下一个可能就大于这个标准值了
            while (end > start && arr[end] >= stanard) end--;
            arr[start] = arr[end];
            while (end > start && arr[start] <= stanard) start++;
            arr[end] = arr[start];
        }
        // 因为上述的交换都少了最开始的start值,因而在这里将其补回
        arr[start] = stanard;
        return start;
    }


}
