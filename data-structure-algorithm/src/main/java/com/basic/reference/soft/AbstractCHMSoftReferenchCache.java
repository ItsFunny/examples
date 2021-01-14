package com.basic.reference.soft;

import com.basic.reference.AbstractReferenceCache;
import com.basic.reference.ClearStrategy;
import com.basic.reference.ObjectCreateStrategy;

import java.lang.ref.ReferenceQueue;
import java.lang.ref.SoftReference;
import java.util.concurrent.ConcurrentHashMap;

/**
 * @author joker
 * @When
 * @Description 高速缓存, 底层是CHM和软引用
 * @Detail 1. 通过ReferenceQueue
 * 2. 获取不到的话就可能是任务
 * @date 创建时间：2019-02-01 06:18
 */
public abstract class AbstractCHMSoftReferenchCache<K, V> extends AbstractReferenceCache<K, V>
{
    protected ClearStrategy<V> clearStrategy;
    protected ReferenceQueue<V> valueQueue;
    protected ConcurrentHashMap<K, SoftReferenceInfo<K, V>> dataMap;

    public AbstractCHMSoftReferenchCache(ObjectCreateStrategy<V> createStrategy)
    {
        super(createStrategy);
        this.dataMap = new ConcurrentHashMap<>();
        this.valueQueue = new ReferenceQueue<>();
    }

    public AbstractCHMSoftReferenchCache(ClearStrategy<V> clearStrategy, ObjectCreateStrategy<V> createStrategy)
    {
        super(createStrategy);
        this.clearStrategy = clearStrategy;
        this.dataMap = new ConcurrentHashMap<>();
        this.valueQueue = new ReferenceQueue<>();

    }


    @Override
    protected void doPut(K key, V value)
    {
        dataMap.put(key, new SoftReferenceInfo<>(key, value, valueQueue));
    }

    @Override
    protected V doGet(K key)
    {
        SoftReferenceInfo<K, V> info;
        info = dataMap.get(key);
        if (null == info)
        {
            return null;
        }
        return info.get();
    }

    @Override
    protected void clear()
    {
        this.clearStrategy.clear(valueQueue);
    }
}
