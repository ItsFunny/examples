package com.basic.reference;

/**
 * @author joker
 * @When
 * @Description 高速缓存
 * 提供了不同的选择:
 * 1.可以基于软引用创建缓存,
 * 2.或者是通过弱引用创建缓存,
 * 底层也提供了不同的选择,
 * 可以通过CHM,也可以通过RingBuffer,
 * 也可以自定义 自定义只需要继承AbstractReferenceCache 实现自定义的底层结构即可
 * @Detail
 * @date 创建时间 ：2019-02-01 06:16
 */

/*
    TODO:
    []  底层采用无锁的RingBuffer


    FIXME:
    建议可以是自定义一个queue,这样的话可以通过长度判断是否要启用分发机制,不然的话每次都锁住太麻烦了(因为clear需要遍历
    整个队列,然后一个一个删除,太慢了),或者是模仿CHM一样,判断节点是什么类型的,
    如果是已经内存回收的节点则clear,否则的话直接创建新的(避免遍历删除)
    避免有的线程get的时候发现为空(原先就没插入过)然后去遍历删除,使得用户等待

 */
public interface IReferenceCache<K, V>
{
    /**
     * 当值不存在的时候会自动添加默认值
     * 并且应当提供一种创建对象的方式,提供一种key-value映射的关系
     * 如涉及业务的情况下:通过key查询对应的值findById等
     * @param key   the key
     * @param value the value
     */
    void put(K key, V value);

    /**
     * Get v.
     *
     * @param key the key
     * @return the v
     */
    V get(K key);


    /***
     * 当值不存在的时候不会再添加
     * @param key the key
     * @return v
     */
    V getIfAbsent(K key);
}

