package com.basic.reference;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-01 07:08
 */
public abstract class AbstractReferenceCache<K, V> implements ReferenceCache<K, V>
{
    protected abstract void doPut(K key, V value);

    protected abstract V doGet(K key);

    protected abstract void clear();

    protected ObjectCreateStrategy<V> createStrategy;


    public AbstractReferenceCache(ObjectCreateStrategy<V> createStrategy)
    {
        this.createStrategy = createStrategy;
    }

    @Override
    public void put(K key, V value)
    {
        doPut(key, value);
    }

    @Override
    public V get(K key)
    {
        V value = doGet(key);
        if (null == value)
        {
            clear();
            value = createStrategy.create(key);
            put(key, value);
        }
        return value;
    }

    @Override
    public V getIfAbsent(K key)
    {
        return doGet(key);
    }
}
