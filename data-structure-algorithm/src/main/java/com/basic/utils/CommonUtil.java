package com.basic.utils;


/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 10:44
 */
public class CommonUtil
{
    public static Integer[] ARR = {15, 23, 1, 0, 16};

    public static void show(Integer[] arr)
    {
        for (Integer integer : arr)
        {
            System.out.printf("%d---", integer);
        }
        System.out.println();
    }

    public static void show()
    {
        show(ARR);
    }
}
