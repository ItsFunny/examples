package com.basic.lock;

import lombok.Data;

import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.TimeUnit;

/**
 * @author joker
 * @When
 * @Description 死锁的实现
 * @Detail 死锁的必要条件:
 * 互斥
 * 占有且等待
 * 循环且等待
 * 不可剥夺
 * @date 创建时间：2019-02-07 13:48
 */
@Data
public class DeadLock
{
    private Object a;
    private Object b;

    public DeadLock()
    {
        this.a=new Object();
        this.b=new Object();
    }

    public void aLock() throws InterruptedException
    {
        synchronized (a)
        {
            // 为了先让其他线程先获取到b锁所以sleep2s
            TimeUnit.SECONDS.sleep(2);
            System.out.println("试图获取到对象b的锁");
            synchronized (b)
            {
                System.out.println("获取到了对象b的锁");
            }
        }
    }

    public void bLock() throws InterruptedException
    {
        synchronized (b)
        {
            // 为了先让其他线程先获取到a锁所以sleep2s
            TimeUnit.SECONDS.sleep(2);
            System.out.println("试图获取到对象a的锁");
            synchronized (a)
            {
                System.out.println("获取到了对象a的锁");
            }
        }
    }

    public static void main(String[] args) throws InterruptedException
    {
        DeadLock lock = new DeadLock();
        new Thread(() ->
        {
            try
            {
                lock.aLock();
            } catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }).start();
        new Thread(() ->
        {
            try
            {
                lock.bLock();
            } catch (InterruptedException e)
            {
                e.printStackTrace();
            }
        }).start();
        TimeUnit.SECONDS.sleep(100);
    }
}
