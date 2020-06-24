package com.basic.sort;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 09:35
 */
public class InsertSortAlgorithm
{
    public static void main(String[] args)
    {
        Integer[] arrs = {14, 25, 23, 1, 25};
        for (Integer arr : arrs)
        {
            System.out.printf("%d---", arr);

        }
        System.out.println();
        InsertSortAlgorithm algorithm = new InsertSortAlgorithm();
        algorithm.insertSort(arrs);

        for (Integer arr : arrs)
        {
            System.out.printf("%d---", arr);

        }
        System.out.println();
    }

    public void insertSort(Integer[] arrs)
    {
        // 插入排序的核心就是假设要插入的元素的下标index 0-index 都是有序的(升序)
        // 所以需要从后往前判断,因为是升序,所以要找到的是坐标是小于他的,而右边是大于他的,因而判断条件是大于
        // 如果条件成立需要将这个值往后移动(不用担心插入的值,因为开始就会保存)
        // 当内层循环跳出的时候也就意味着,下标所在的值是小于temp值的,而我们需要在这index+1处插入
        // 因为之前的数都已经往后移动了
        for (int i = 1; i < arrs.length; i++)
        {
            int temp = arrs[i];
            int j = i - 1;
            for (; j >= 0 && arrs[j] >= temp; j--)
            {
                arrs[j + 1] = arrs[j];
            }
            arrs[j + 1] = temp;

        }
    }
}