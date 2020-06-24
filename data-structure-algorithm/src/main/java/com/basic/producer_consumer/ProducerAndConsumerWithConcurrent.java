package com.basic.producer_consumer;

import java.util.UUID;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

/**
 * @author joker
 * @When
 * @Description 生产者消费者的实现
 * @Detail 10个消费者, 5个生产者, 通过concurrent组件
 * @date 创建时间：2019-02-07 14:03
 */
public class ProducerAndConsumerWithConcurrent
{
    // 内部是链表,并且有读锁和写锁
    private LinkedBlockingQueue<String> foods;

    public ProducerAndConsumerWithConcurrent()
    {
        this.foods = new LinkedBlockingQueue<>();
    }

    final Runnable PRODUCER = () ->
    {
        while (!Thread.currentThread().isInterrupted())
        {
            try
            {
                TimeUnit.SECONDS.sleep(1);
                String food = UUID.randomUUID().toString();
                foods.put(food);
                System.out.println("生产者:" + Thread.currentThread().getName() + " 新增食物: " + food);
            } catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }
    };
    final Runnable CONSUMER = () ->
    {
        while (!Thread.currentThread().isInterrupted())
        {
            try
            {
                String food = foods.take();
                System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                TimeUnit.MILLISECONDS.sleep(800);
            } catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }
    };

    public void test()
    {
        // 10个生产者
        for (int i = 0; i < 10; i++)
        {
            new Thread(PRODUCER).start();
        }
        // 5个消费者
        for (int i = 0; i < 5; i++)
        {
            new Thread(CONSUMER).start();
        }
    }

    public static void main(String[] args) throws InterruptedException
    {
        new ProducerAndConsumerWithConcurrent().test();
        while (true)
        {
            TimeUnit.SECONDS.sleep(100);
        }
    }
}
