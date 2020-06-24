package com.basic.structure.list;

import sun.misc.Cache;

import javax.xml.crypto.Data;
import java.util.ArrayList;
import java.util.Iterator;

/**
 * @author joker
 * @When
 * @Description LRU 最近最久未使用,
 * @Detail 大体类似于HashMap, 但是不同的地方在于
 * 也是最核心的地方在于,会将数据都链接起来
 * @date 创建时间：2019-02-17 18:21
 */
public class LruCache implements Iterable
{
    @Override
    public Iterator iterator()
    {
        return new CacheIterator();
    }

    private class CacheNode
    {
        int key;
        int data;
        int count;
        CacheNode previous;
        CacheNode next;

        public CacheNode(int data)
        {
            this.data = data;
        }

        @Override
        public String toString()
        {
            return "CacheNode{" +
                    "key=" + key +
                    ", data=" + data +
                    '}';
        }
    }

    private class CacheIterator implements Iterator
    {
        CacheNode node;

        public CacheIterator()
        {
            this.node = head;
        }

        @Override
        public boolean hasNext()
        {
            return this.node != null;
        }

        @Override
        public Object next()
        {
            Object data = node.data;
            node = node.next;
            return data;
        }
    }

    CacheNode head, tail;
    int capility;
    // fixedsize
    CacheNode[] nodes;
    int size;

    public LruCache(int capility)
    {
        // capility必须为2的n次幂
        this.capility = capility;
        this.nodes = new CacheNode[capility];
    }

    /**
     * Add.
     * 便于方便,这里的key,value都为int类型
     *
     * @param key   the key
     * @param value the value
     */
    public void add(int key, int value)
    {
        int index = key & (this.capility - 1);
        CacheNode[] nods = this.nodes;
        CacheNode node = null;
        CacheNode newNode = new CacheNode(value);
        if ((node = nods[index]) == null)
        {
            nods[index] = newNode;
        } else
        {
            CacheNode keyNode = null;
            // 锁住
//            synchronized (node)
//            {
            while (node != null)
            {
                if (node.key == key)
                {
                    keyNode = node;
                    break;
                }
                node = node.next;
            }
//            }
            if (null != keyNode)
            {
//              synchronized (keyNode)
                keyNode.count++;
                // 是否需要重新排序待定
                // 注意如果使用lru的话就不适合使用synchronized了,而应该使用lock
                // 但在这里我们假设单线程使用
                // 这里的话直接将这个元素放到尾部

                // 接下来的步骤就是重新连接了,将这个节点的前一个节点指向之后,并且需要判断是否是头节点和尾节点以及是否是数组的首元素
                if (keyNode == nodes[index])
                {
                    // 数组的头结点
                    nodes[index] = node.next;
                    //
                } else if (keyNode == this.head)
                {   // 头节点
                    this.head = keyNode.next;
                } else if (keyNode == this.tail)
                {
                    // 尾节点,如果是尾节点则啥也不用干
                    return;
                } else
                {
                    // 则这个节点的上一个节点需要连接到后一个节点,后一个节点的prev需要连接到前一个
                    keyNode.previous.next = keyNode.next;
                    // 已经不是尾节点了,所以不可能为空
                    keyNode.next.previous = keyNode.previous;
                }
                this.tail.next = keyNode;
                keyNode.next = null;
                this.tail = keyNode;
                return;
            }
        }
        // 插入
        if (this.size == 0)
        {
            this.head = newNode;
            this.size++;
        } else if (this.size + 1 > this.capility)
        {
            // 移除表头元素
            this.head = this.head.next;
            this.tail.next = newNode;
        } else
        {
            this.tail.next = newNode;
            this.size++;
        }
        this.tail = newNode;
    }

    public void show()
    {
        CacheNode cacheNode = this.head;
        while (null != cacheNode)
        {
            System.out.println(cacheNode);
            cacheNode = cacheNode.next;
        }
    }
}
