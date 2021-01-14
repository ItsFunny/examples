package com.basic.structure.list;

import lombok.Data;

import java.lang.reflect.Array;
import java.util.Iterator;
import java.util.*;
import java.util.function.Consumer;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-20 12:14
 */


/*
    ArrayList中的遍历是通过内部的Itar,内部中有个cursor属性来达到遍历的效果
 */
@Data
public class MyArrayList<T> implements Iterable<T>
{
    T[] data;
    int maxSize;
    int index;
    int threshold;
    Class type;

    public MyArrayList(Class type, int maxSize, float loadFactory)
    {
        if (loadFactory >= 1 || loadFactory <= 0) loadFactory = 0.75f;
        this.type = type;
        this.maxSize = maxSize;
        this.data = (T[]) Array.newInstance(type, maxSize);
        this.threshold = (int) (maxSize * loadFactory);
    }

    public void add(T value)
    {
        if (null == data)
        {
            throw new NullPointerException();
        }
        if (index + 1 >= threshold) this.resize();
        data[index++] = value;
    }

    public T remove()
    {
        T value = data[index];
        data[index--] = null;
        return value;
    }


    public void resize()
    {
        T[] values = this.data;
        T[] newData = (T[]) Array.newInstance(this.type, maxSize << 1);
//        System.arraycopy();
        for (int i = 0; i < data.length; i++)
        {
            newData[i] = values[i];
        }
        this.data = newData;
        Map<String, String> map = new HashMap<>();

    }

    private class MyIterator implements Iterator<T>
    {
        int cursor;

        @Override
        public boolean hasNext()
        {
            return cursor < index;
        }

        @Override
        public T next()
        {
            return data[cursor++];
        }
    }

    @Override
    public Iterator<T> iterator()
    {
        return new MyIterator();
    }

}
