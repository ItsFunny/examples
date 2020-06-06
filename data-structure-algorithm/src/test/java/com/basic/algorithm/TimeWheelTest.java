package com.basic.algorithm;

import com.basic.algorithm.timewheel.TimeInterface;
import com.basic.algorithm.timewheel.TimeWheel;
import lombok.Data;
import org.junit.Test;

import java.util.*;
import java.util.concurrent.*;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-24 23:18
 */
public class TimeWheelTest
{
    @Data
    static class TestClass implements TimeInterface
    {
        public static Integer staticId = 0;
        private final Integer id = staticId++;

        @Override
        public void callBack()
        {
            System.out.println("class:" + id + "call callback");
        }

        @Override
        public Integer getKey()
        {
            Random random = new Random();
            return random.nextInt(100);
        }
    }

    @Test
    public void testTimeWheel() throws InterruptedException
    {
        TimeWheel timeWheel = new TimeWheel(10, TimeUnit.SECONDS, 1l);
        for (int i = 0; i < 20; i++)
        {
            timeWheel.addDelayJob(new TestClass());
        }
        new Thread(() ->
        {
            timeWheel.start();
        }).start();

        TimeUnit.SECONDS.sleep(200);
        TreeSet<String> treeSet = new TreeSet<>();
    }

}
