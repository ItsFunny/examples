package com.basic.structure.queue;

import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-20 16:03
 */
public class StackImplQueueTest
{

    @Test
    public void testStackImplQueue()
    {
        TwoStackImplQueue<Integer> twoStackImplQueue = new TwoStackImplQueue<>();
        List<Integer> list = Arrays.asList(1, 2, 3, 4, 5, 6);
        for (Integer integer : list)
        {
            twoStackImplQueue.push(integer);
        }
        List<Integer> resultList = new ArrayList<>();
        Integer temp = twoStackImplQueue.pop();
        while (temp != null)
        {
            resultList.add(temp);
            temp = twoStackImplQueue.pop();
        }
        for (Integer integer : resultList)
        {
            System.out.printf("%d->",integer);
        }
        assert resultList.equals(list);
    }

}
