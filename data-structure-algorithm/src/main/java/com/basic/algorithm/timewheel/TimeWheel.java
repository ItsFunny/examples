package com.basic.algorithm.timewheel;

import lombok.Data;
import lombok.extern.slf4j.Slf4j;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * @author joker
 * @When
 * @Description 时间轮算法的实现
 * 模拟延迟队列的实现
 * @Detail
 * @date 创建时间：2019-01-24 16:52
 */
/*
    内部的数据结构为了防止hash冲突,采用的是链地址法
    时间轮每隔一个duraiton转动一次,这次采取游标的形式来代替时间轮的转动
    V1 不提供运行时数组扩容
    V2 分层时间轮
    V3 采用RingBuffer代替CopyOnWriteArrayList(无锁CAS的形式)
    V3 消费任务分发到线程池


 */
@Data
@Slf4j
// @ThreadNotSafe
public class TimeWheel
{


    // 环状缓存
    private CopyOnWriteArrayList<Object> slots;

    // 游标
    private int index = -1;

    // slot长度最长为多少
    private Integer threashold;

    private AtomicInteger size;

    private boolean stopFlag;

    // 时间间隔
    private TimeUnit tickyType;
    private Long sleepDuration;


    @Data
    public static class SlotNode
    {
        private TimeInterface service;
        private SlotNode next;
        private SlotNode tail;

        public SlotNode(TimeInterface service)
        {
            this.service = service;
        }

        public SlotNode()
        {
        }

        public void setService(TimeInterface service)
        {
            this.service = service;
        }

        public void insertNode(TimeInterface service)
        {
            SlotNode newSlotNode = new SlotNode(service);
            this.tail.next = newSlotNode;
            this.tail = newSlotNode;
        }
    }

    @Data
    public static class Slot
    {
        private SlotNode root;

        public Slot(SlotNode root)
        {
            this.root = root;
        }
    }

    public TimeWheel(Integer threashold, TimeUnit tickyType, Long sleepDuration)
    {
        this.threashold = threashold;
        Object[] objects = new Object[this.threashold];
        this.slots = new CopyOnWriteArrayList<>(objects);
        this.tickyType = tickyType;
        this.sleepDuration = sleepDuration;
        this.size = new AtomicInteger(0);

    }

    public void addDelayJob(TimeInterface service)
    {
        log.info("[addDelayJob]receive one job:{}", service);
        Integer index = hash(service.getKey());
        Integer size = this.slots.size();
        // 说明slots还没满,则先将slots搞满
        if (index > size)
        {
            // ex: index=6 ,size=3 ,insert 4,5 with null then add 4,5 first
            // TODO 优化
            List<?> tempList = new ArrayList<>();
            for (int i = 0; i < index - size - 1; i++)
            {
                this.size.incrementAndGet();
                tempList.add(null);
            }
            this.slots.addAll(tempList);
            this.slots.add(new Slot(new SlotNode(service)));
        } else
        {
            Object o = this.slots.get(index);
            SlotNode slotNode = null;
            if (null == o)
            {
                Slot slot = new Slot(new SlotNode(service));
                this.slots.set(index, slot);
                slot.root.tail = slot.root;
            } else
            {
                Slot slot = (Slot) o;
                slot.root.insertNode(service);
            }
        }
        this.size.incrementAndGet();
    }

    public void executeJob(Integer index)
    {
        Object o = this.slots.get(index);
        if (null == o)
        {
            return;
        }
        Slot slot = (Slot) o;
        SlotNode tempNode = slot.root;
        // 遍历执行
        while (null != tempNode)
        {
            tempNode.getService().callBack();
            tempNode = tempNode.getNext();
            slot.root = tempNode;
        }

    }

    public void start()
    {
        // 每隔1s转动一次
        long currentTimemills = System.currentTimeMillis();
        while (!this.stopFlag)
        {
//            while (currentTimemills < currentTimemills + 1000) ;
//            currentTimemills = System.currentTimeMillis();
            try
            {
                this.tickyType.sleep(this.sleepDuration);
            } catch (InterruptedException e)
            {
                log.error("[start]error:{}", e);
                this.stopFlag = true;
                // clear all
            }

            if (++this.index >= this.threashold)
            {
                this.index = 0;
            }
            this.executeJob(this.index);

            System.out.println("执行完一个任务,还有"+this.size.decrementAndGet()+"个任务");
        }
    }


    protected Integer hash(Integer key)
    {
        return key % this.threashold;
    }

}
