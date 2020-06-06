package com.basic.lock;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author joker
 * @When
 * @Description 读写锁
 * @Detail
 * @date 创建时间：2019-02-08 11:03
 */
public class ReadWriteLock
{
    Object writeObj, readObj;
    List<String> foods;


    public ReadWriteLock()
    {
        this.writeObj = new Object();
        this.readObj = new Object();
        this.foods = new ArrayList<>();
    }

    final Runnable WRITER = () ->
    {
        // 读写锁,写的时候是排它锁
        // 也就意味着需要先抢占到读

    };
    final Runnable READER = () ->
    {
        // 读的时候共享锁
        // 意味着读的时候不能写,所以要确保
        synchronized (writeObj)
        {

        }
    };
}
