package com.basic.producer_consumer;

import java.util.LinkedList;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-07 16:36
 */
public class ProducerAndConsumerWIthLockAndCondition
{
    private Lock lock;
    private Condition takeCondition;
    private Condition putCondition;
    private List<String> foods;

    public ProducerAndConsumerWIthLockAndCondition()
    {
        this.lock = new ReentrantLock();
        this.takeCondition = this.lock.newCondition();
        this.putCondition = this.lock.newCondition();
        this.foods = new LinkedList<>();
    }

    final Runnable PRODUCER = () ->
    {
        lock.lock();
        try
        {

            while (!Thread.currentThread().isInterrupted())
            {
                // 如果引入了putCondition
                // 就需要判断容量来限制了
                while (this.foods.size() == 16) this.putCondition.await();
                String food = UUID.randomUUID().toString();
                foods.add(food);
                takeCondition.signalAll();
            }

        } catch (Exception e)
        {
            e.printStackTrace();
        } finally
        {
            lock.unlock();
        }
        try
        {
            TimeUnit.SECONDS.sleep(1);
        } catch (InterruptedException e)
        {
            e.printStackTrace();
        }
    };
    final Runnable CONSUMER = () ->
    {
        lock.lock();
        try
        {
            while (!Thread.currentThread().isInterrupted())
            {
                while (foods.isEmpty()) takeCondition.await();
                String food = foods.remove(0);
                System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                putCondition.signalAll();
            }
        } catch (Exception e)
        {
            e.printStackTrace();
        } finally
        {
            lock.unlock();
        }
        try
        {
            TimeUnit.MILLISECONDS.sleep(800);
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
        new ProducerAndConsumerWIthLockAndCondition().test();
    }
}
