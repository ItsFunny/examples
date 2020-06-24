package com.basic.producer_consumer;

import lombok.Data;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

/**
 * @author joker
 * @When
 * @Description 通过Object自带的notify/wait/waitAll实现生产者消费者模型
 * @Detail
 * @date 创建时间：2019-02-07 15:41
 */
@Data
public class ProducerAndConsumerWithNotify
{
    private List<String> foods;

    public ProducerAndConsumerWithNotify()
    {
        this.foods = new LinkedList<>();
    }

    final Runnable PRODUCER = () ->
    {
        try
        {
            while (!Thread.currentThread().isInterrupted())
            {
                String food = UUID.randomUUID().toString();
                synchronized (foods)
                {
//                    while (foods.size() == 16)
//                    {
//                        // 如果放不下了,则就需要进行等待了
//                        foods.wait();
//                    }
                    foods.add(food);
                    System.out.println("生产者:" + Thread.currentThread().getName() + " 新增食物: " + food);
                    // 提示所有的消费者都可以消费了
                    foods.notifyAll();
                }
                TimeUnit.SECONDS.sleep(1);
            }
        } catch (InterruptedException e)
        {
            e.printStackTrace();
        }
    };

    final Runnable CONSUMER = () ->
    {
        try
        {
            while (!Thread.currentThread().isInterrupted())
            {
                synchronized (foods)
                {
                    // 对于消费者而言,如果foods为空,代表着需要等待producer生产食物
                    // 注意这里必须是while循环
                    // 因为当线程重新被唤醒之后,因为程序计数器,从而会继续在这里执行
                    // 而如果producer生产速度慢,当1号consumer消费完毕,2号抢到了之后也会notifyall,
                    // 如果恰巧是consumer获取到了则会跳出了if循环,从而直接remove(0),也就报index错误了,while则会继续先判断
                    while (foods.isEmpty()) foods.wait();
                    String food = foods.remove(0);
                    System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                    foods.notifyAll();
                }
                TimeUnit.MILLISECONDS.sleep(800);
            }
        } catch (InterruptedException e)
        {
            e.printStackTrace();
        }
    };

    public void test()
    {
        for (int i = 0; i < 10; i++)
        {
            new Thread(CONSUMER).start();
        }
        for (int i = 0; i < 5; i++)
        {
            new Thread(PRODUCER).start();
        }

        while (true) ;
    }

    public static void main(String[] args)
    {
        new ProducerAndConsumerWithNotify().test();
    }
}
