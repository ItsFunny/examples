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
 * @date 创建时间：2019-01-20 17:31
 */
public class TwoQueueImplStackTest
{

    @Test
    public void testTwoQueueImplStack()
    {
        TwoQueueImplStack<Integer> twoQueueImplStack = new TwoQueueImplStack<>();
        List<Integer> integerList = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9);
        for (Integer integer : integerList)
        {
            twoQueueImplStack.push(integer);
        }
        Integer temp = twoQueueImplStack.pop();
        List<Integer> resultList = new ArrayList<>();
        while (null != temp)
        {
            resultList.add(temp);
            temp = twoQueueImplStack.pop();
        }
        for (Integer integer : resultList)
        {
            System.out.printf("%d->", integer);
        }
        assert resultList.equals(Arrays.asList(9, 8, 7, 6, 5, 4, 3, 2, 1));
    }


}
