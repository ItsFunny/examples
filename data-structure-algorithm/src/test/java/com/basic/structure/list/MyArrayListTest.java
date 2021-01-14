package com.basic.structure.list;

import org.junit.Test;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-20 14:52
 */
public class MyArrayListTest
{

    @Test
    public void testMyArrayList()
    {
        MyArrayList<Integer> myArrayList = new MyArrayList<Integer>(Integer.class, 10, 0.75f);
        myArrayList.add(0);
        myArrayList.add(1);
        myArrayList.add(2);
        myArrayList.add(3);
        myArrayList.add(4);
        for (Integer integer : myArrayList)
        {
            System.out.println(integer);
        }
        // 删除元素
        myArrayList.remove();
        System.out.println("=========");
        for (Integer integer : myArrayList)
        {
            System.out.println(integer);
        }
    }

}
