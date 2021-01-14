package stack;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Charlie
 * @When
 * @Description 用队列实现栈
 * @Detail 时刻保持一个队列为空, 另外一个队列有数据即可
 * @Attention: 遇到一个坑, 就是当遍历的时候, 采用的是q1.size ,但是当poll的时候,却忽略了size会变
 * @Date 创建时间：2020-03-16 16:18
 */
public class Stack_225_Implement_Stack_using_Queues
{
    Queue<Integer> q1;
    Queue<Integer> q2;

    /**
     * Initialize your data structure here.
     */
    public Stack_225_Implement_Stack_using_Queues()
    {
        q1 = new LinkedList<>();
        q2 = new LinkedList<>();

    }

    /**
     * Push element x onto stack.
     */
    public void push(int x)
    {
        if (q1.isEmpty() && q2.isEmpty())
        {
            q1.add(x);
        } else if (!q1.isEmpty())
        {
            q1.add(x);
        } else
        {
            q2.add(x);
        }
    }

    /**
     * Removes the element on top of the stack and returns that element.
     */
    public int pop()
    {

        if (!q1.isEmpty())
        {
            int size = q1.size();
            // 将元素移动到q2
            for (int i = 0; i < size - 1; i++)
            {
                q2.add(q1.poll());
            }
            return q1.poll();
        } else
        {
            int size = q2.size();
            for (int i = 0; i < size - 1; i++)
            {
                q1.add(q2.poll());
            }
            return q2.poll();
        }
    }

    /**
     * Get the top element.
     */
    public int top()
    {
        if (!q1.isEmpty())
        {
            // 将元素移动到q2
            int size = q1.size();
            for (int i = 0; i < size - 1; i++)
            {
                q2.add(q1.poll());
            }
            Integer peek = q1.poll();
            q2.add(peek);
            return peek;
        } else
        {
            int size = q2.size();
            for (int i = 0; i < size - 1; i++)
            {
                q1.add(q2.poll());
            }
            Integer peek = q2.poll();
            q1.add(peek);
            return peek;
        }
    }

    /**
     * Returns whether the stack is empty.
     */
    public boolean empty()
    {
        return q1.isEmpty() && q2.isEmpty();
    }

    public static void main(String[] args)
    {

        Stack_225_Implement_Stack_using_Queues q = new Stack_225_Implement_Stack_using_Queues();
        q.push(1);
        q.push(2);
        q.push(3);
        System.out.println(q.pop());
        System.out.println(q.pop());
        System.out.println(q.pop());
        System.out.println(q.empty());
    }
}

