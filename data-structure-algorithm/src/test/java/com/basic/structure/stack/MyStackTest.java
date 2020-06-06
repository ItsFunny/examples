package com.basic.structure.stack;

import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-20 16:40
 */
public class MyStackTest
{


    @Test
    public void testMyStack()
    {
        MyStack<Integer> myStack = new MyStack<>();
        List<Integer> startList = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8);
        for (Integer integer : startList)
        {
            myStack.push(integer);
        }
        List<Integer> resultList = new ArrayList<>();
        Integer temp = myStack.pop();
        while (null != temp)
        {
            resultList.add(temp);
            temp = myStack.pop();
        }
        for (Integer integer : resultList)
        {
            System.out.printf("%d->", integer);
        }
        assert resultList.equals(Arrays.asList(8, 7, 6, 5, 4, 3, 2, 1));

    }
}
