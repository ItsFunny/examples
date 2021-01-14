package com.basic.structure.tree;

import java.util.concurrent.TimeUnit;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-19 15:38
 */
public class ThreadLocalTest
{
    static ThreadLocal<String> threadLocal = new ThreadLocal<>();

    public static void main(String[] args)
    {
        for (Integer i = 0; i < 2; i++)
        {
            final Integer t = i;
            new Thread(new Runnable()
            {
                @Override
                public void run()
                {
                    threadLocal.set("q" + t);
                    try
                    {
                        TimeUnit.SECONDS.sleep(2);
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                    System.out.println(threadLocal.get());
                }
            }).start();
        }
    }

}
