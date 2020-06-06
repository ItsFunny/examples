package com.basic.structure.stack;

import lombok.Data;

import java.util.ArrayList;
import java.util.Stack;
/**
 * @author joker
 * @When
 * @Description 原生栈的简单实现, 因为上面的list数组采用的是泛型, 所以这里的数组直接采用object
 * 栈是先进后出的数据结构,而在我们的实际实现中需要每次添加元素都添加到末尾,这样pop的时候就不需要拷贝数组了
 * @Detail
 * @date 创建时间：2019-01-20 16:13
 */
@Data
public class MyStack<T>
{

    Object[] data;
    int cursor;
    int maxSize;

    public MyStack()
    {
        this.maxSize = Integer.MAX_VALUE;
        this.data = new Object[100];
    }

    public void push(T value)
    {
        if (cursor + 1 >= maxSize)
        {
            // 扩容
            throw new IndexOutOfBoundsException();
        }
        data[cursor++] = value;
    }

    public T pop()
    {
        if (cursor <= 0)
        {
            cursor = 0;
            return null;
        }
        T value=(T) data[--cursor];
        data[cursor]=null;
        return value;
    }

}
