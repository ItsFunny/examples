package com.basic.reference;

import com.basic.reference.soft.CHMSoftReferenceCache;
import org.junit.Test;

import java.util.Random;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-01 11:05
 */
public class SoftReferenceCacheTest
{
    // 如何测试呢: 软引用是当内存不足,先测试一下minor gc是否会回收
    // 1. 测试minor gc 是否会触发回收这些内存 -XX -XX:MaxNewSize=30M   ->minor gc 并不会回收
    @Test
    public void testSoftReferenceCache()
    {
        // 创建20m的内存,限制年轻代内存为30m
        ReferenceCache<String, byte[]> cache = new CHMSoftReferenceCache<>((key) -> new byte[1024 * 1024 * 4 * 5]);
        byte[] bytes = cache.get("1");
        bytes = null; // 手动解引用
        cache.get("2");
        cache.get("1");
    }

    // 2. 测试full gc 是否会回收   -Xmx40M 促使发生GC
    // 结论: 当full gc 前才会回收 ,对象会被回收到referenceQueue中,然后那个对象会被置空
    @Test
    public void testSoftReferenceCacheFullGC()
    {
        ReferenceCache<String, byte[]> cache = new CHMSoftReferenceCache<>((key) -> new byte[1024 * 1024 * 4 * 5]);
        byte[] bytes = cache.get("1");
        bytes = null; // 手动解引用
        cache.get("2");
        byte[] bytes1 = cache.get("1");
        System.out.println(bytes1.length);
    }


}
