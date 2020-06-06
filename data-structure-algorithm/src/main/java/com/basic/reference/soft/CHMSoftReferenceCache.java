package com.basic.reference.soft;

import com.basic.reference.ClearStrategy;
import com.basic.reference.ObjectCreateStrategy;
import com.basic.reference.ReferenceCache;


/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-01 06:42
 */
public class CHMSoftReferenceCache<K, V> extends AbstractCHMSoftReferenchCache<K, V>
{
    private ClearStrategy<V> DEFAULT_CHM_CLEAR_STRATEGY = (queue) ->
    {
        SoftReferenceInfo<K,V> poll = (SoftReferenceInfo<K, V>) queue.poll();
        System.out.println(poll);
        while (null != poll)
        {
            K k=poll.getKey();
            System.out.println("clear key="+k);
            this.dataMap.remove(k);
        }
    };


    public CHMSoftReferenceCache(ClearStrategy<V> clearStrategy, ObjectCreateStrategy<V> objectCreateStrategy)
    {
        super(clearStrategy, objectCreateStrategy);
    }

    public CHMSoftReferenceCache(ObjectCreateStrategy<V> objectCreateStrategy)
    {
        super(objectCreateStrategy);
        this.clearStrategy = DEFAULT_CHM_CLEAR_STRATEGY;
    }

    public static void main(String[] args)
    {
        ReferenceCache<String, byte[]> cache = new CHMSoftReferenceCache<>((key) -> new byte[1024 * 1024 * 4 * 5]);
        byte[] bytes = cache.get("1");
        bytes = null; // 手动解引用
        System.gc();
        cache.get("2");
        byte[] bytes1 = cache.get("1");
        System.out.println(bytes1.length);
    }

}
