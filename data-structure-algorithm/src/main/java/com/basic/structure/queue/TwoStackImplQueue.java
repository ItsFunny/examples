package com.basic.structure.queue;

import lombok.Data;

import java.util.Stack;
import java.util.concurrent.Semaphore;
import java.util.concurrent.locks.AbstractQueuedSynchronizer;


/**
 * @author joker
 * @When
 * @Description 两个栈实现一个队列
 * @Detail 队列是先进先出的数据结构, 而栈则是先进后出的数据结构, 如何通过2个栈实现数据的先进先出呢
 * 逻辑如下: 当数据入栈push的时候 将数据压入栈A,试图出栈达到先进先出的时候,将栈A中的数据先pop到栈B中,这样
 * 顺序就相反了,如: A中:a->b->c->d->e  当出栈压入B的时候: e->d->c->b->a 这样a就相当于后进的了
 * @date 创建时间：2019-01-20 15:42
 */
@Data
public class TwoStackImplQueue<T>
{
    // 名字待定
    private Stack<T> stack1;
    private Stack<T> stack2;

    public TwoStackImplQueue()
    {
        this.stack1 = new Stack<>();
        this.stack2 = new Stack<>();
    }

    public void push(T data)
    {
        this.stack1.push(data);
    }

    public T pop()
    {
        // 如果stack2不为空,则先将其内部的元素弹出去(这时候已经是先进先出的了)
        if (!stack2.isEmpty())
        {
            return stack2.pop();
        }
        // 安全校验
        if (stack1.isEmpty())
        {
            return null;
        }
        // 这里可以省去一部操作的,如注释所示,可以省去一个入栈
        while (!stack1.isEmpty())
        {
            stack2.push(stack1.pop());
        }
        // 弹出最后一个
//        for (int i = 0; i < stack1.size()-1; i++)
//        {
//            stack2.push(stack1.pop());
//        }
//        return stack1.pop();
        return stack2.pop();
    }


    public static void main(String[] args)
    {
        Semaphore semaphore=new Semaphore(5);
    }

}
