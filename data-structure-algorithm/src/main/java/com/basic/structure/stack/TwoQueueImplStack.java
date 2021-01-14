package com.basic.structure.stack;

import lombok.Data;

import java.util.LinkedList;
import java.util.concurrent.ConcurrentHashMap;

/**
 * @author joker
 * @When
 * @Description 两个队列实现一个栈
 * 队列是先进先出,而栈是先进后出,如何做到通过2和队列实现先进后出
 * 逻辑其实与2个栈实现一个队列是类似的,一个队列A用于push数据,而另外一个队列B用于pop数据,也相同在pop之前先将
 * 队列A的数据pop直至剩下一个: 如原先a->b->c->d 最终只会剩下d,而a->b->c则会到队列B中
 * 然后下一次push的时候会往不为空的队列中push,所以总结要点就是:
 * 1. 2个队列,push,pop操作,肯定是一个队列为空,另外一个队列不为空
 * 2. push的时候讲元素添加到不为空的队列中
 * 3. pop的时候,不为空的队列将数据除了最后一个外(这个最后一个就是要求的值)全都pop然后push到另外一个队列中,
 * @Detail
 * @date 创建时间：2019-01-20 17:03
 */
@Data
public class TwoQueueImplStack<T>
{
    private LinkedList<T> queue1;
    private LinkedList<T> queue2;

    public TwoQueueImplStack()
    {
        this.queue1 = new LinkedList<>();
        this.queue2 = new LinkedList<>();
    }

    // 注意这里的linkedlist是要用add的,add是添加到末尾,而push是添加到第一个
    // 而我们要的pop是从末尾弹出的
    // push时候的注意点,我们要保持一个队列为空,另外一个队列不为空,所以我们每次添加都是往有值的队列中添加元素
    // 至于第一次添加,就随意了,给第一个还是第二个都可以
    public void push(T value)
    {
        if (queue1.isEmpty())
        {
            queue2.add(value);
        }else if (!queue2.isEmpty())
        {
            throw new RuntimeException("逻辑错误,内部某块逻辑错误");
        }else{
            queue1.add(value);
        }
    }

    // 必须保持2个队列一个为空,另外一个不为空
    // 出队就是将一个队列中的元素移到另外一个队列
    public T pop()
    {
        T temp = null;
        if (!queue1.isEmpty() && queue2.isEmpty())
        {
            // 最后一个元素是我们想要的元素(既最后入队的最先弹出)
            for (int i = 0; i < queue1.size() - 1; i++)
            {
                temp = queue1.pop();
                queue2.add(temp);
            }
            temp = queue1.pop();
        } else if (!queue2.isEmpty() && queue1.isEmpty())
        {
            // 最后一个元素是我们想要的元素(既最后入队的最先弹出)
            for (int i = 0; i < queue2.size() - 1; i++)
            {
                temp = queue2.pop();
                queue2.add(temp);
            }
            temp = queue2.pop();
        } else if (!queue1.isEmpty() )
        {
            // 说明2个队列都不为空,这是不可能也不该出现的
            throw new RuntimeException("逻辑错误");
        } else
        {
            return null;
        }
        return temp;
    }
}
