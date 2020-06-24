package com.basic.structure.list;

import org.junit.Test;

import java.util.Iterator;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-17 21:09
 */
public class LruCacheTest
{
    // lru 简单版测试
    @Test
    public void testLru()
    {
        LruCache cache = new LruCache(16);
        for (int i = 0; i < 16; i++)
        {
            cache.add(i, i);
        }
        for (Object o : cache)
        {
            System.out.println(o);
        }
        System.out.println("++++++++++++");
//        cache.show();
        cache.add(22, 22);

//        while (iterator.hasNext())
//        {
//            Object next = iterator.next();
//            System.out.println(next);
//            iterator= (Iterator) iterator.next();
//        }
        for (Object o : cache)
        {
            System.out.println(o);
        }
        System.out.println("===========");
//        cache.show();
        cache.add(9, 9);
        for (Object o : cache)
        {
            System.out.println(o);
        }
//        cache.show();
    }
}
